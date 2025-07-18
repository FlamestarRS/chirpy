package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("forbidden: dev environment only"))
		return
	}
	cfg.fileserverHits.Store(0)

	err := cfg.db.ResetUsers(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error resetting users database" + err.Error()))
	}

	err = cfg.db.ResetRefreshTokens(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error resetting refresh tokens database" + err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0, databasses successfully reset"))
}
