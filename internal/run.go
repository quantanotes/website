package internal

import (
	"log/slog"
	"net/http"
	_ "quanta/internal/handlers"
	"quanta/internal/single"
)

func Run() {
	slog.Info("listening on localhost:3000")
	err := http.ListenAndServe(":3000", single.Router)
	slog.Error(err.Error())
}
