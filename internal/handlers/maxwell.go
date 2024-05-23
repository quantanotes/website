package handlers

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
)

var maxwellRoots = protJSONResHandler(
	func(ctx context.Context, user string) (model.Knowledge, error) {
		return model.GetOrCreateRootThread(ctx, user)
	},
)

var maxwellChildren = protJSONResHandler(
	func(ctx context.Context, user string) ([]model.Knowledge, error) {
		return model.GetThreadChildren(ctx, user)
	},
)

func maxwellThreads(w http.ResponseWriter, r *http.Request) {

}

func maxwellCreate(w http.ResponseWriter, r *http.Request) {

}

func maxwellDelete(w http.ResponseWriter, r *http.Request) {

}

func maxwellChat(w http.ResponseWriter, r *http.Request) {
	res, err := http.Post(globals.MaxwellURL+"/chat", "application/json", r.Body)
	if err != nil {
		errorInternalResponse(w, err)
		return
	}
	defer res.Body.Close()

	w.Header().Set("content-type", "application/json")

	_, err = io.Copy(w, res.Body)
	if err != nil {
		slog.Error(err.Error())
	}
}
