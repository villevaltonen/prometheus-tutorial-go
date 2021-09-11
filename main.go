package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = prometheus.NewCounter(prometheus.CounterOpts{Name: "hello_world_counter"})

func main() {
	prometheus.MustRegister(counter)

	r := mux.NewRouter()
	r.HandleFunc("/api/hello", hello()).GetMethods()
	r.Handle("/metrics", promhttp.Handler()).GetMethods()

	log.Println("Starting the app...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}

func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world!")
		counter.Inc()
		w.Write([]byte(fmt.Sprintf("OK")))
	}
}