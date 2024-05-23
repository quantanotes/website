package handlers

import (
	"context"
	"quanta/internal/model"
)

var grant = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct {
		Id         string
		Permission int
	}) (int, error) {
		return model.GrantLinkPermissionWhereAuthor(ctx, req.Id, user, req.Permission)
	},
)

var retrieveLinks = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct {
		Id string
	}) ([]model.Knowledge, error) {
		return model.GetLinksWhereAuthor(ctx, req.Id, user)
	},
)

var createLink = protJSONReqJSONResHandler(
	func(ctx context.Context, user string, req struct {
		Id string
	}) (model.Knowledge, error) {
		return model.CreateLinkWhereAuthor(ctx, req.Id, user)
	},
)

var deleteLink = protJSONReqOKResHandler(
	func(ctx context.Context, user string, req struct {
		Id string
	}) error {
		return model.DeleteLinkWhereAuthor(ctx, req.Id, user)
	},
)
