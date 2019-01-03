package jwt

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var jwks map[string]*rsa.PublicKey

// CreateParser - initialize JWT parser
// if 'keys' is nil, JWKS will be fetched from os.Getenv("JWKS_URL") via HTTP
func CreateParser(keys map[string]*rsa.PublicKey) *Parser {
	if keys != nil {
		jwks = keys
		return new(Parser)
	}
	jwks = FetchJWKS()
	return new(Parser)
}

// Parser - JWT parser
type Parser struct {
	UseJSONNumber        bool // Use JSON Number format in JSON decoder
	SkipClaimsValidation bool // Skip claims validation during token parsing
}

type mapClaims map[string]interface{}

func (m mapClaims) Valid() error {
	return nil
}

// Parse - Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
// If everything is kosher, err will be nil
func (p *Parser) Parse(tokenString string) (*Token, error) {
	return p.ParseWithClaims(tokenString, mapClaims{})
}

// ParseWithClaims - parse, validate, and return a token with claims
func (p *Parser) ParseWithClaims(tokenString string, claims Claims) (*Token, error) {
	token, parts, err := p.ParseUnverified(tokenString, claims)
	if err != nil {
		return token, err
	}

	// Verify signing method is correct - only RS256 is supported
	if token.Header["alg"] != "RS256" {
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
	if err = verify(strings.Join(parts[0:2], "."), token.Signature, key); err != nil {
		return nil, err
	}

	token.Valid = true
	return token, nil
}

// ParseUnverified - parse token without verifying signature
// WARNING: Don't use this method unless you know what you're doing
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
	// JSON Decode. Special case for map type to avoid weird pointer behavior
	if c, ok := token.Claims.(mapClaims); ok {
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

func verify(signingString string, signature string, publicKey *rsa.PublicKey) error {
	var err error

	// // Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write([]byte(signingString))

	// Verify the signature
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hasher.Sum(nil), sig)
}
