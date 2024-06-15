package internal

import (
	"log/slog"
	"net/http"
	"quanta/internal/handlers"
)

func Run() {
	slog.Info("Listening on localhost:3000")
	err := http.ListenAndServe(":3000", handlers.R)
	slog.Error(err.Error())
}
