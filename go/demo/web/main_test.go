package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handlerFunc(res, req)
	exp := "Hello World"
	var dat map[string]interface{}
	jerr := json.Unmarshal(res.Body.Bytes(), &dat)
	if jerr != nil {
		t.Fatal(jerr)
	}
	if exp != dat["message"] {
		t.Fatalf("Expected %s gog %s", exp, dat["message"])
	}
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(Web())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	exp := "Hello World"
	if exp != string(body) {
		t.Fatalf("Expected %s got %s", exp, body)
	}
}
