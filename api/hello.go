package handler

import (
	"net/http"

	"github.com/luispfcanales/lambda-vercel/pkg/response"
)

// Hello is handler
func Hello(w http.ResponseWriter, r *http.Request) {
	response.SendJSON(w, "send reponse to request /hello")
}
