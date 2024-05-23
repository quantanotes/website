package handlers

import (
	"context"
	"quanta/internal/model"
)

var spaceAlgorithm = pubJSONResHandler(
	func(ctx context.Context) (any, error) {
		return model.GetKnowledgeWherePublic(ctx)
	},
)

var spaceChildren = pubJSONReqJSONResHandler(
	func(ctx context.Context, req struct{ Id string }) ([]model.Knowledge, error) {
		return model.GetKnowledgeChildren(ctx, req.Id)
	},
)
