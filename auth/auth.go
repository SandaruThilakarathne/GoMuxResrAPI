package auth

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Msg string `json:"msg"`
}

func Test(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{"hello"})
}
