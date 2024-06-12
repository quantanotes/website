package handlers

import (
	"context"
	"quanta/internal/globals"
	"quanta/internal/model"
	"quanta/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

var rootTitle = "root"

// Unused - saving for future
// var nodeRoots = protectedJSONResponse(func(ctx context.Context, userID string) ([]model.Node, error) {
// 	return model.GetNodeChildrenWhereAuthor(ctx, userID, userID)
// })

var nodeRoot = protectedJSONResponse(func(ctx context.Context, userID string) (model.Node, error) {
	category := chi.URLParamFromCtx(ctx, "category")
	// If we are accessing root data for Notes or Io
	if category != "note" && category != "thread" {
		return model.Node{}, globals.ErrBadRequest("Invalid or missing category.")
	}

	node, err := model.GetNodeChildWhereAuthorCategory(ctx, userID, userID, category)
	if err == pgx.ErrNoRows {
		node, err = model.CreateNodeWhereAuthor(ctx, model.Node{
			Parent: userID, Author: userID, Category: category, Title: &rootTitle,
		})
	}

	return node, err
})

type nodeChildrenRequest struct {
	ID string
}

var nodeChildren = protectedJSONRequestJSONResponse(func(ctx context.Context, userID string, req nodeChildrenRequest) ([]model.Node, error) {
	return model.GetNodeChildrenWhereAuthor(ctx, req.ID, userID)
})

var nodeCreate = protectedJSONRequest(func(ctx context.Context, userID string, req model.Node) error {
	req.Author = userID
	utils.ReplaceNilWithEmptyString(&req.Title)
	utils.ReplaceNilWithEmptyString(&req.Location)
	utils.ReplaceNilWithEmptyString(&req.Content)
	_, err := model.CreateNodeWhereAuthor(ctx, req)
	return err
})

var nodeUpdate = protectedJSONRequest(func(ctx context.Context, userID string, req model.Node) error {
	req.Author = userID
	return model.UpdateNodeWhereAuthor(ctx, req)
})

type nodeDeleteRequest struct {
	ID string
}

var nodeDelete = protectedJSONRequest(func(ctx context.Context, userID string, req nodeDeleteRequest) error {
	return model.DeleteNodeWhereAuthor(ctx, req.ID, userID)
})

type nodeMoveRequest struct {
	ID     string
	Parent string
}

var nodeMove = protectedJSONRequest(func(ctx context.Context, userID string, req nodeMoveRequest) error {
	return model.MoveNodeWhereAuthor(ctx, req.ID, userID, req.ID)
})
