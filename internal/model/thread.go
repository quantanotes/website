package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	threadCategory  = "thread"
	threadRootTitle = "root"
)

func CreateThread(ctx context.Context, parent string, author string) (Knowledge, error) {
	return CreateKnowledge(ctx, parent, author, threadCategory, "", "")
}

func UpdateThread(ctx context.Context, id string, author string, title string, content string) error {
	return UpdateKnowledgeWhereAuthor(ctx, id, author, title, content)
}

func DeleteThread(ctx context.Context, id string, author string) error {
	return DeleteKnowledgeWhereAuthor(ctx, id, author)
}

func MoveThread(ctx context.Context, id string, author string, to string) error {
	return MoveKnowledgeWhereAuthor(ctx, id, author, to)
}

func GetOrCreateRootThread(ctx context.Context, author string) (Knowledge, error) {
	thread, err := GetKnowledgeChildWhereAuthorCategory(ctx, author, author, threadCategory)
	if err == pgx.ErrNoRows {
		thread, err = CreateKnowledge(ctx, author, author, threadCategory, threadRootTitle, "")
	}
	return thread, err
}

func GetThreadChildren(ctx context.Context, parent string, author string) ([]Knowledge, error) {
	return GetKnowledgeChildrenWhereAuthor(ctx, parent, author)
}

func PublishThread(ctx context.Context, id string, author string) (int, error) {
	return GrantPublicReadPermissionWhereAuthor(ctx, id, author)
}

func UnpublishThread(ctx context.Context, id string, author string) (int, error) {
	return RevokePublicReadPermissionWhereAuthor(ctx, id, author)
}
