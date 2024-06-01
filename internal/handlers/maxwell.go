package handlers

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
)

var maxwellRoots = protJSONResHandler(
	func(ctx context.Context, user string) (model.Knowledge, error) {
		return model.GetOrCreateRootThread(ctx, user)
	},
)

var maxwellChildren = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) ([]model.Knowledge, error) {
		return model.GetThreadChildren(ctx, req.Id, user)
	},
)

var maxwellCreate = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct{ Parent string }) error {
		_, err := model.CreateThread(ctx, req.Parent, user)
		return err
	},
)

var maxwellDelete = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) error {
		return model.DeleteThread(ctx, req.Id, user)
	},
)

var maxwellUpdate = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct {
		Id      string
		Title   string
		Content string
	}) error {
		return model.UpdateThread(ctx, req.Id, user, req.Title, req.Content)
	},
)

var maxwellMove = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct {
		Id string
		To string
	}) error {
		return model.MoveThread(ctx, req.Id, user, req.To)
	},
)

var maxwellPublish = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) (int, error) {
		return model.UnpublishThread(ctx, req.Id, user)
	},
)

var maxwellUnpublish = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) (int, error) {
		return model.UnpublishThread(ctx, req.Id, user)
	},
)

func maxwellChat(w http.ResponseWriter, r *http.Request) {
	res, err := http.Post(globals.MaxwellURL+"/chat", "application/json", r.Body)
	if err != nil {
		errorInternalResponse(w, err)
		return
	}
	defer res.Body.Close()

	w.Header().Set("content-type", "application/json")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// Write the line to the response writer
		_, err := w.Write([]byte(line + "\n"))
		if err != nil {
			log.Println("Error writing to response:", err)
			return
		}
		flusher.Flush()
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from response body:", err)
	}
}
