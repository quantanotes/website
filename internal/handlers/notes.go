package handlers

import (
	"context"
	"quanta/internal/model"
)

var notesRoots = protJSONResHandler(
	func(ctx context.Context, user string) (model.Knowledge, error) {
		return model.GetOrCreateRootNote(ctx, user)
	},
)

var notesCreate = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct{ Parent string }) error {
		_, err := model.CreateNote(ctx, req.Parent, user)
		return err
	},
)

var notesDelete = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) error {
		return model.DeleteNote(ctx, req.Id, user)
	},
)

var notesUpdate = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct {
		Id      string
		Title   string
		Content string
	}) error {
		return model.UpdateNote(ctx, req.Id, user, req.Title, req.Content)
	},
)

var notesMove = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct {
		Id string
		To string
	}) error {
		return model.MoveNote(ctx, req.Id, user, req.To)
	},
)

var notesChildren = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) ([]model.Knowledge, error) {
		return model.GetNoteChildren(ctx, req.Id, user)
	},
)

var notesPublish = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) (int, error) {
		return model.PublishNote(ctx, req.Id, user)
	},
)

var notesUnpublish = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct{ Id string }) (int, error) {
		return model.UnpublishNote(ctx, req.Id, user)
	},
)
