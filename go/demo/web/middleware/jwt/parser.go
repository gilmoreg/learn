package jwt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

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
func (p *Parser) Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
	return p.ParseWithClaims(tokenString, MapClaims{}, keyFunc)
}

func (p *Parser) ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
	token, parts, err := p.ParseUnverified(tokenString, claims)
	if err != nil {
		return token, err
	}

	// Verify signing method is correct
	if token.Header["alg"] != "RS256" {
		// signing method is not in the listed set
		return token, fmt.Errorf("signing method %v is invalid", token.Header["alg"])
	}

	// Lookup key
	var key interface{}
	if keyFunc == nil {
		// keyFunc was not provided.  short circuiting validation
		return token, fmt.Errorf("no Keyfunc was provided")
	}
	if key, err = keyFunc(token); err != nil {
		return token, err
	}

	// Validate Claims
	if !p.SkipClaimsValidation {
		if err := token.Claims.Valid(); err != nil {
			return token, err
		}
	}

	// Perform validation
	token.Signature = parts[2]
	if err = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, key); err != nil {
		vErr.Inner = err
		vErr.Errors |= ValidationErrorSignatureInvalid
	}

	if vErr.valid() {
		token.Valid = true
		return token, nil
	}

	return token, vErr
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
