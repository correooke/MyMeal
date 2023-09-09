package handlers

import "net/http"

func IsAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive!"))
}
