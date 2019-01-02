package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gilmoreg/learn/go/demo/web/models"
)

// Hello controller
func Hello(w http.ResponseWriter, r *http.Request) {
	message := models.HelloMessage{Message: "Hello World"}
	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
