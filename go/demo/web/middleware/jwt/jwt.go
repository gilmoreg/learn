package jwt

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWK struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	X5c []string `json:"x5c"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
}

type JWKSResponse struct {
	Keys []JWK `json:"keys"`
}

var jwks map[string]*rsa.PublicKey

func init() {
	fmt.Println("Fetching JWKS")
	resp, err := http.Get(os.Getenv("JWKS_URL"))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var jwksResponse JWKSResponse
	err = json.Unmarshal(body, &jwksResponse)
	if err != nil {
		panic(err)
	}

	for _, key := range jwksResponse.Keys {
		var publicKey, err = createPublicKey(key)
		if err != nil {
			panic(err)
		}
		jwks[key.Kid] = publicKey
	}

	fmt.Println("Fetched JWKS")
}

type SigningMethodRSA struct {
	Name string
	Hash crypto.Hash
}

func createPublicKey(key JWK) (*rsa.PublicKey, error) {
	nStr := key.N
	eStr := key.E

	// Base64URL Decode the strings
	decN, _ := base64.RawURLEncoding.DecodeString(nStr)
	decE, _ := base64.RawURLEncoding.DecodeString(eStr)

	n := big.NewInt(0)
	n.SetBytes(decN)

	// Pad decE if it is less than 8 bytes.
	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}

	eReader := bytes.NewReader(eBytes)
	var e uint64
	err := binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	publicKey := rsa.PublicKey{N: n, E: int(e)}

	return &publicKey, nil
}

func (m *SigningMethodRSA) Verify(signingString, signature string, key interface{}) error {
	var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
		return err
	}

	var rsaKey *rsa.PublicKey
	var ok bool

	if rsaKey, ok = key.(*rsa.PublicKey); !ok {
		return errors.New("invalid key type")
	}

	// Create hasher
	if !m.Hash.Available() {
		return errors.New("hash unavailable")
	}
	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	// Verify the signature
	return rsa.VerifyPKCS1v15(rsaKey, m.Hash, hasher.Sum(nil), sig)
}

func getTokenFromRequest(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) == 2 && auth[0] == "Bearer" {
		return auth[1]
	}
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return accessToken.Value
}

func getSigningKey(token *jwt.Token) (*rsa.PublicKey, error) {
	kid := token.Header["kid"].(string)
	fmt.Println(kid)
	key, ok := jwks[kid]
	if !ok {
		return nil, errors.New("No matching key")
	}
	return key, nil
}

// VerifyJWT - middleware verifying JWT signature
func VerifyJWT(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("verifying JWT")
	tokenString := getTokenFromRequest(r)
	fmt.Println(tokenString)
	if tokenString == "" {
		http.Error(rw, "Not authorized", 401)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return getSigningKey(token)
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		next(rw, r)
	} else {
		fmt.Println(err)
		http.Error(rw, "Not authorized", 401)
		return
	}
}
