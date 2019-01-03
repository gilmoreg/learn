package jwt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func Verify(signingString string, signature string, publicKey *rsa.PublicKey) error {
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
