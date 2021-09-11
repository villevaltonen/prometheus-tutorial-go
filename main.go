package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	prom_metrics "github.com/deathowl/go-metrics-prometheus"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcrowley/go-metrics"
)

var counter prometheus.Counter = prometheus.NewCounter(prometheus.CounterOpts{Name: "hello_world_counter"})

func main() {
	client := prom_metrics.NewPrometheusProvider(metrics.DefaultRegistry, "", "", prometheus.DefaultRegisterer, 1*time.Second)
	prometheus.MustRegister(counter)
	go client.UpdatePrometheusMetrics()

	r := mux.NewRouter()
	r.HandleFunc("/api/hello", hello()).GetMethods()
	r.Handle("/metrics", promhttp.Handler()).GetMethods()

	log.Println("Starting the app...")
	http.ListenAndServe(":8080", r)
}

func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		w.Write([]byte(fmt.Sprintf("OK")))
	}
}