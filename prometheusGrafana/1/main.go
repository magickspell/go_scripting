package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

//запускаем доекер и смотрим что у нас в логах го:
//http://localhost:9090/metrics

//просто порты прометеуса и графаны
//http://localhost:9090
//http://localhost:3000/login

func main() {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}
