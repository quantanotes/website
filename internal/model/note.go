package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	noteCategory  = "note"
	noteRootTitle = "root"
)

func CreateNote(ctx context.Context, parent string, author string) (Knowledge, error) {
	return CreateKnowledge(ctx, parent, author, noteCategory, "", "")
}

func UpdateNote(ctx context.Context, id string, author string, title string, content string) error {
	return UpdateKnowledgeWhereAuthor(ctx, id, author, title, content)
}

func DeleteNote(ctx context.Context, id string, author string) error {
	return DeleteKnowledgeWhereAuthor(ctx, id, author)
}

func MoveNote(ctx context.Context, id string, author string, to string) error {
	return MoveKnowledgeWhereAuthor(ctx, id, author, to)
}

func GetOrCreateRootNote(ctx context.Context, author string) (Knowledge, error) {
	note, err := GetKnowledgeChildWhereAuthorCategory(ctx, author, author, noteCategory)
	if err == pgx.ErrNoRows {
		note, err = CreateKnowledge(ctx, author, author, noteCategory, noteRootTitle, "")
	}
	return note, err
}

func GetNoteChildren(ctx context.Context, parent string, author string) ([]Knowledge, error) {
	return GetKnowledgeChildrenWhereAuthorCategory(ctx, parent, author, noteCategory)
}

func PublishNote(ctx context.Context, id string, author string) (int, error) {
	return GrantPublicReadPermissionWhereAuthor(ctx, id, author)
}

func UnpublishNote(ctx context.Context, id string, author string) (int, error) {
	return RevokePublicReadPermissionWhereAuthor(ctx, id, author)
}
