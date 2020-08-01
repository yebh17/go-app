package main

import (
	"fmt"
	"net/http"

	chiprometheus "github.com/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	n := chi.NewRouter()
	m := chiprometheus.NewMiddleware("api", 50, 100, 200, 500, 1000, 5000)
	// if you want to use other buckets than the default (300, 1200, 5000) you can run:
	// m := negroniprometheus.NewMiddleware("serviceName", 400, 1600, 700)

	n.Use(m)

	n.Handle("/metrics", promhttp.Handler())
	n.Get("/name", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "This is Sunny\nGlad to meet you bruh!")
	})
	n.Get("/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hey man! wassup ;-)")
	})
	n.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 Not Found")
	})

	http.ListenAndServe(":6777", n)
}
