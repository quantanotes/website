package handlers

import (
	"context"
	"quanta/internal/model"
)

var spaceRoots = jsonResponseHandler(model.GetNodesWherePublic)

type spaceChildrenRequest struct {
	Parent string
}

var spaceChildren = jsonRequestJSONResponseHandler(func(ctx context.Context, req spaceChildrenRequest) ([]model.Node, error) {
	return model.GetNodeChildrenWherePublic(ctx, req.Parent)
})
