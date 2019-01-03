package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const keysize = 2048

func GenerateKeyPair() (*rsa.PrivateKey, error) {
	rng := rand.Reader
	keypair, err := rsa.GenerateKey(rng, keysize)
	if err != nil {
		return nil, err
	}

	return keypair, nil
}

func MessageSignature(message []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privatekey, crypto.SHA256, hashed[:])
	if err != nil {
		return []byte{}, err
	}

	return signature, nil
}

type JWTPayload map[string]interface{}

type JWTHeader map[string]interface{}

func CreateJWT(payload JWTPayload, privatekey *rsa.PrivateKey) ([]byte, error) {
	header := map[string]interface{}{
		"alg": "RS256",
		"typ": "JWT",
		"kid": "test",
	}
	headerSerialized, err := json.Marshal(header)
	if err != nil {
		return []byte{}, err
	}
	stringheader := base64.RawURLEncoding.EncodeToString(headerSerialized)

	payloadSerialized, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, err
	}
	stringpayload := base64.RawURLEncoding.EncodeToString(payloadSerialized)

	presignedsig := stringheader + "." + stringpayload
	signature, err := MessageSignature([]byte(presignedsig), privatekey)
	if err != nil {
		return []byte{}, err
	}

	base64sig := base64.RawURLEncoding.EncodeToString(signature)
	return []byte(presignedsig + "." + base64sig), nil
}

// CreateJWTWithNewKey is just a sample function to easily create a JWT, you would
// never want to expose a private key in this way
func CreateJWTWithNewKey(payload JWTPayload) ([]byte, *rsa.PrivateKey, error) {
	key, err := GenerateKeyPair()
	if err != nil {
		return []byte{}, nil, err
	}

	jwt, err := CreateJWT(payload, key)
	if err != nil {
		return []byte{}, nil, err
	}

	return jwt, key, nil
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatal(fmt.Sprintf("%v != %v", a, b))
	}
}

func TestMissingTokenShouldNotPass(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, req, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized)
}

func TestInvalidAccessTokenCookieShouldNotPass(t *testing.T) {
	w := httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: "garbage"})
	request := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, request, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized)
}

func TestInvalidBearerTokenShouldNotPass(t *testing.T) {
	w := httptest.NewRecorder()
	request := &http.Request{Header: http.Header{"Authorization": []string{"Bearer garbage"}}}
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, request, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized)
}

func TestValidAcessTokenCookieShouldPass(t *testing.T) {
	w := httptest.NewRecorder()
	payload := mapClaims{"test": "test"}
	token, key, err := CreateJWTWithNewKey(JWTPayload{"test": "test"})
	if err != nil {
		t.Fatalf("Error creating token and key")
	}
	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: string(token)})
	request := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &key.PublicKey})
	var called bool
	middleware(w, request, func(w http.ResponseWriter, r *http.Request) {
		called = true
		val := r.Context().Value(Key("user")).(mapClaims)
		assertEqual(t, reflect.DeepEqual(val, payload), true)
	})
	assertEqual(t, called, true)
}
