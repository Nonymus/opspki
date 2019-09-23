package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	toldYouSo = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "told_you_so_events_total",
			Help: "Total number of Told You So events",
		},
		[]string{"user"},
	)
	listenAddress = flag.String("listenAddress", ":2121", "Listen address")
)

func main() {
	flag.Parse()
	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/tys/{user}", ToldYouSo).Methods("POST")
	r.HandleFunc("/tys/{user}", NotAllowed).Methods("GET", "OPTIONS", "PUT", "DELETE")

	log.Printf("Starting webserver on \"%s\"", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, r))
}

func NotAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
}

func ToldYouSo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	toldYouSo.With(prometheus.Labels{"user": user}).Inc()
}
