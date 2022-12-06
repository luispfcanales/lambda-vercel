package api

import (
	"fmt"
	"net/http"
)

// Home is handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}