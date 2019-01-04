package jwt

import (
	"context"
	"crypto/rsa"
	"errors"
	"net/http"
	"strings"
)

var parser *Parser

type Key string

func getTokenFromRequest(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) == 2 && auth[0] == "Bearer" {
		return auth[1]
	}
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}
	return accessToken.Value
}

func getSigningKey(token *Token) (*rsa.PublicKey, error) {
	kid := token.Header["kid"].(string)
	key, ok := jwks[kid]
	if !ok {
		return nil, errors.New("No matching key")
	}
	return key, nil
}

// VerifyJWT - middleware verifying JWT signature
func verifyJWT(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenString := getTokenFromRequest(r)
	if tokenString == "" {
		http.Error(rw, "Not authorized", 401)
		return
	}

	token, err := parser.Parse(tokenString)
	if err != nil {
		http.Error(rw, "Not authorized", 401)
		return
	}

	if _, ok := token.Claims.(mapClaims); ok && token.Valid {
		ctx := r.Context()
		ctx = context.WithValue(ctx, Key("user"), token.Claims.(mapClaims))
		next(rw, r.WithContext(ctx))
	} else {
		http.Error(rw, "Not authorized", 401)
		return
	}
}

// VerifyJWT - factory function for JWT validation middleware
func VerifyJWT(keys map[string]*rsa.PublicKey) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	parser = CreateParser(keys)
	return verifyJWT
}
