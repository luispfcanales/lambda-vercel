package handler

import (
	"fmt"
	"net/http"
)

// Hello is handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
