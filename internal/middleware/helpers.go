package middleware

import (
	"log/slog"
	"net/http"
	"quanta/internal/globals"
)

// Exact same function as in handlers.
// CBA to do module gymnastics just for 1 function so we copy it
func writeError(w http.ResponseWriter, err error) {
	switch err := err.(type) {
	case globals.Error:
		w.WriteHeader(err.Status)
		w.Write([]byte(err.Message))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened on our end. Please try again later."))
		slog.Error(err.Error())
	}
}
