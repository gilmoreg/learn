package jwt

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestMissingTokenShouldNotPass(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, req, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized, "")
}

func TestInvalidAccessTokenCookieShouldNotPass(t *testing.T) {
	w := httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: "garbage"})
	request := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, request, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized, "")
}

func TestInvalidBearerTokenShouldNotPass(t *testing.T) {
	w := httptest.NewRecorder()
	request := &http.Request{Header: http.Header{"Authorization": []string{"Bearer garbage"}}}
	middleware := VerifyJWT(map[string]*rsa.PublicKey{"test": &rsa.PublicKey{}})
	middleware(w, request, func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Should not reach here")
	})
	assertEqual(t, w.Result().StatusCode, http.StatusUnauthorized, "")
}
