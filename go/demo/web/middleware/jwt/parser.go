package jwt

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

type Parser struct {
	UseJSONNumber        bool // Use JSON Number format in JSON decoder
	SkipClaimsValidation bool // Skip claims validation during token parsing
}

type MapClaims map[string]interface{}

func (m MapClaims) Valid() error {
	return nil
}

// Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
// If everything is kosher, err will be nil
func (p *Parser) Parse(tokenString string) (*Token, error) {
	return p.ParseWithClaims(tokenString, MapClaims{})
}

func (p *Parser) ParseWithClaims(tokenString string, claims Claims) (*Token, error) {
	token, parts, err := p.ParseUnverified(tokenString, claims)
	if err != nil {
		return token, err
	}

	// Verify signing method is correct
	if token.Header["alg"] != "RS256" {
		// signing method is not in the listed set
		return token, fmt.Errorf("signing method %v is invalid", token.Header["alg"])
	}

	// Validate Claims
	if !p.SkipClaimsValidation {
		if err := token.Claims.Valid(); err != nil {
			return token, err
		}
	}

	// Lookup key
	key, ok := jwks[token.Header["kid"].(string)]
	if !ok {
		return nil, errors.New("no matching key")
	}

	// Perform validation
	token.Signature = parts[2]
	if err = Verify(strings.Join(parts[0:2], "."), token.Signature, key); err != nil {
		return nil, err
	}

	token.Valid = true
	return token, nil
}

// WARNING: Don't use this method unless you know what you're doing
//
// This method parses the token but doesn't validate the signature. It's only
// ever useful in cases where you know the signature is valid (because it has
// been checked previously in the stack) and you want to extract values from
// it.
func (p *Parser) ParseUnverified(tokenString string, claims Claims) (token *Token, parts []string, err error) {
	parts = strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, parts, errors.New("token contains an invalid number of segments")
	}

	token = &Token{Raw: tokenString}

	// parse Header
	var headerBytes []byte
	if headerBytes, err = DecodeSegment(parts[0]); err != nil {
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			return token, parts, errors.New("tokenstring should not contain 'bearer'")
		}
		return token, parts, err
	}
	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, parts, err
	}

	// parse Claims
	var claimBytes []byte
	token.Claims = claims

	if claimBytes, err = DecodeSegment(parts[1]); err != nil {
		return token, parts, err
	}
	dec := json.NewDecoder(bytes.NewBuffer(claimBytes))
	if p.UseJSONNumber {
		dec.UseNumber()
	}
	// JSON Decode.  Special case for map type to avoid weird pointer behavior
	if c, ok := token.Claims.(MapClaims); ok {
		err = dec.Decode(&c)
	} else {
		err = dec.Decode(&claims)
	}
	// Handle decode error
	if err != nil {
		return token, parts, err
	}

	return token, parts, nil
}
