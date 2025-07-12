package main

import (
	"html/template"
	"net/http"
)

type data struct {
	Hits int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("metrics.html")
	if err != nil {
		http.Error(w, "error loading template", http.StatusInternalServerError)
		return
	}

	data := data{
		Hits: cfg.fileserverHits.Load(),
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}
