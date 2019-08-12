package handler

import (
	"fmt"
	"net/http"
)

// Handler is the main handler for logging in using the twitter API
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
