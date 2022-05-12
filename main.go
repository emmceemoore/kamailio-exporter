package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		listenAddress               = ":9494"
		metricsPath                 = "/metrics"
		scrapeUri                   = "unix:/var/run/kamailio/kamailio_ctl"
		timeout       time.Duration = 5 * time.Second
	)

	collector, err := NewCollector(scrapeUri, timeout)

	if err != nil {
		panic(err)
	}

	prometheus.MustRegister(collector)

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body><a href="` + metricsPath + `">metrics</a></body></html>`))
	})
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
