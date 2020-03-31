package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	toldYouSo = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "told_you_so_events_total",
			Help: "Total number of Told You So events",
		},
	)
	listenAddress = flag.String("listenAddress", ":2121", "Listen address")
)

func main() {
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/toldyouso", ToldYouSoHandler)

	log.Printf("Starting webserver on \"%s\"", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func ToldYouSoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} else {
		toldYouSo.Inc()
	}
}