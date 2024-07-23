package main

import "net/http"

func (cfg *config) handlerReadiness(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	respondWithJSON(w, http.StatusOK, payload)
}
