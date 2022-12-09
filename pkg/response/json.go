package response

import (
	"encoding/json"
	"net/http"
)

// ReponseHTTP is model to response request
type ReponseHTTP struct {
	Status  uint        `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SendJSON write in format JSON
func SendJSON(w http.ResponseWriter, txt string) {
	w.Header().Add("content-type", "application/json")
	rs := &ReponseHTTP{
		Status:  200,
		Message: txt,
	}
	json.NewEncoder(w).Encode(rs)
}
