package utils

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	StatusCode int
	Data       interface{}
}

func SendRes(w http.ResponseWriter, r *http.Request, res *Res) {
	requestID := r.Context().Value("request-id")
	if requestID != nil {
		w.Header().Set("Request-ID", requestID.(string))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	resByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Write(resByte)
	return
}
