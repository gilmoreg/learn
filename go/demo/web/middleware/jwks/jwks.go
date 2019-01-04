package jwks

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type jwk struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	X5c []string `json:"x5c"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
}

type jwksResponse struct {
	Keys []jwk `json:"keys"`
}

func createPublicKey(key jwk) (*rsa.PublicKey, error) {
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

// FetchJWKS - Fetch JWKS from URL
func FetchJWKS(url string) map[string]*rsa.PublicKey {
	fmt.Println("Fetching JWKS")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var jwksResponse jwksResponse
	err = json.Unmarshal(body, &jwksResponse)
	if err != nil {
		panic(err)
	}

	var jwks map[string]*rsa.PublicKey
	for _, key := range jwksResponse.Keys {
		var publicKey, err = createPublicKey(key)
		if err != nil {
			panic(err)
		}
		jwks[key.Kid] = publicKey
	}

	fmt.Println("Fetched JWKS")
	return jwks
}
