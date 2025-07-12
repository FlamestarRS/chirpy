package main

import (
	"net/http"
	"sync/atomic"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	apiCfg := apiConfig{}

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidate)

	server.ListenAndServe()

}

type apiConfig struct {
	fileserverHits atomic.Int32
}
