package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	threadCategory  = "thread"
	threadRootTitle = "root"
)

func GetOrCreateRootThread(ctx context.Context, author string) (Knowledge, error) {
	thread, err := GetKnowledgeChildWhereAuthorCategory(ctx, author, author, threadCategory)
	if err == pgx.ErrNoRows {
		thread, err = CreateKnowledge(ctx, author, author, threadCategory, threadRootTitle, "")
	}
	return thread, err
}

func GetThreadChildren(ctx context.Context, author string) ([]Knowledge, error) {
	return GetKnowledgeChildrenWhereAuthorCategory(ctx, author, author, threadCategory)
}
