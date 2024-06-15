package handlers

import (
	"context"
	"quanta/internal/globals"
	"quanta/internal/model"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

var rootTitle = "root"

// Unused - saving for future
// var nodeRoots = protectedJSONResponse(func(ctx context.Context, userID string) ([]model.Node, error) {
// 	return model.GetNodeChildrenWhereAuthor(ctx, userID, userID)
// })

var nodeRoot = protectedJSONResponseHandler(func(ctx context.Context, userID string) (model.Node, error) {
	category := chi.URLParamFromCtx(ctx, "category")
	// If we are accessing root data for Notes or Io
	if category != "note" && category != "thread" {
		return model.Node{}, globals.ErrBadRequest("Invalid or missing category.")
	}

	node, err := model.GetNodeChildWhereAuthorCategory(ctx, userID, userID, category)
	if err == pgx.ErrNoRows {
		node = model.Node{Parent: userID, Author: userID, Category: category, Title: &rootTitle}
		node.ReplaceNilWithEmpty()
		node, err = model.CreateNodeWhereAuthor(ctx, node)
	}

	return node, err
})

type nodeChildrenRequest struct {
	ID string
}

var nodeChildren = protectedJSONRequestJSONResponseHandler(func(ctx context.Context, userID string, req nodeChildrenRequest) ([]model.Node, error) {
	return model.GetNodeChildrenWhereAuthor(ctx, req.ID, userID)
})

var nodeCreate = protectedJSONRequestHandler(func(ctx context.Context, userID string, req model.Node) error {
	req.Author = userID
	req.ReplaceNilWithEmpty()
	_, err := model.CreateNodeWhereAuthor(ctx, req)
	return err
})

var nodeUpdate = protectedJSONRequestHandler(func(ctx context.Context, userID string, req model.Node) error {
	req.Author = userID
	return model.UpdateNodeWhereAuthor(ctx, req)
})

type nodeDeleteRequest struct {
	ID string
}

var nodeDelete = protectedJSONRequestHandler(func(ctx context.Context, userID string, req nodeDeleteRequest) error {
	return model.DeleteNodeWhereAuthor(ctx, req.ID, userID)
})

type nodeMoveRequest struct {
	ID     string
	Parent string
}

var nodeMove = protectedJSONRequestHandler(func(ctx context.Context, userID string, req nodeMoveRequest) error {
	return model.MoveNodeWhereAuthor(ctx, req.ID, userID, req.ID)
})

type nodePublishRequest struct {
	ID string
}

var nodePublish = protectedJSONRequestHandler(func(ctx context.Context, userID string, req nodePublishRequest) error {
	return model.UpdateNodePermissionWhereAuthor(ctx, req.ID, userID, 1)
})

var nodeUnpublish = protectedJSONRequestHandler(func(ctx context.Context, userID string, req nodePublishRequest) error {
	return model.UpdateNodePermissionWhereAuthor(ctx, req.ID, userID, 0)
})
