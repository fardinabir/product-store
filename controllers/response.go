package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func (r *Response) JSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Body)
	return
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
