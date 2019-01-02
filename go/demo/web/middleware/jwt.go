package middleware

import (
	"bytes"
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

type JWKS struct {
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
	Keys []JWKS `json:"keys"`
}

var jwks JWKSResponse

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
	err = json.Unmarshal(body, &jwks)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fetched JWKS")
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
	kid := token.Header["kid"]
	fmt.Println(kid)
	for _, key := range jwks.Keys {
		fmt.Println(key.N)
		if key.Kid == kid {
			// Extract n and e values from the jwk
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
	}
	return nil, errors.New("No matching key")
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
