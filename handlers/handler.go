package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	for k, v := range r.Header {
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, currentTime)
}
