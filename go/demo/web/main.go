package main

import (
	"fmt"
	"net/http"

	"github.com/gilmoreg/learn/go/demo/web/controllers"
	"github.com/gilmoreg/learn/go/demo/web/middleware/jwt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
	negroniprometheus "github.com/zbindenren/negroni-prometheus"
)

func main() {
	fmt.Println("Starting server on port 3000")

	n := negroni.New()
	m := negroniprometheus.NewMiddleware("web")
	n.Use(m)
	n.Use(negroni.HandlerFunc(middleware.VerifyJWT))

	r := mux.NewRouter()
	r.Handle("/metrics", prometheus.Handler())
	r.HandleFunc("/", controllers.Hello)
	n.UseHandler(r)

	http.ListenAndServe(":3000", n)
}
