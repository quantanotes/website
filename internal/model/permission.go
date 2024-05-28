package model

import (
	"context"
	"quanta/internal/single"
)

const (
	linkCategory = "link"
)

func GrantPublicReadPermissionWhereAuthor(ctx context.Context, id string, author string) (int, error) {
	permission := 0
	query := `
		UPDATE knowledge
		SET permission = 1
		WHERE permission < 1 AND id = $1 AND author_id = $2
		RETURNING permission
	`
	err := single.DB.QueryRow(ctx, query, id, author).Scan(&permission)
	return permission, err
}

func RevokePublicReadPermissionWhereAuthor(ctx context.Context, id string, author string) (int, error) {
	permission := 0
	query := `
		UPDATE knowledge
		SET permission = 0
		WHERE id = $1 AND author_id = $2
		RETURNING permission
	`
	err := single.DB.QueryRow(ctx, query, id, author).Scan(&permission)
	return permission, err
}

func GetLinksWhereAuthor(ctx context.Context, parent string, author string) ([]Knowledge, error) {
	return GetKnowledgeChildrenWhereAuthorCategory(ctx, parent, author, linkCategory)
}

func CreateLinkWhereAuthor(ctx context.Context, parent string, author string) (Knowledge, error) {
	return CreateKnowledgeWhereParentAuthor(ctx, parent, author, linkCategory, "", "", 1)
}

func DeleteLinkWhereAuthor(ctx context.Context, id string, author string) error {
	return DeleteKnowledgeWhereAuthor(ctx, id, author)
}

func GrantLinkPermissionWhereAuthor(ctx context.Context, id string, author string, permission int) (int, error) {
	query := `
		UPDATE knowledge
		SET permission = $3
		WHERE permission < 1 AND id = $1 AND author_id = $2
		RETURNING permission
	`
	err := single.DB.QueryRow(ctx, query, id, author, permission).Scan(&permission)
	return permission, err
}

// func GrantPublicReadPermissionWhereAuthor(ctx context.Context, id string, author string) error {
// 	query := `
// 		INSERT INTO reference (from_id, to_id, category)
// 		SELECT $1, '00000000-0000-0000-0000-000000000000', 'read'
// 		WHERE EXISTS (
// 			SELECT 1
// 			FROM knowledge
// 			WHERE
// 				id = $1
// 				AND author_id = $2
// 				AND category = 'read'
// 		)
// 		AND NOT EXISTS (
// 			SELECT 1
// 			FROM reference
// 			WHERE
// 				from_id = $1
// 				AND to_id = '00000000-0000-0000-0000-000000000000'
// 				AND category = 'read'
// 		)
// 	`

// 	_, err := single.DB.Exec(ctx, query, id, author)
// 	return err
// }

// func RevokePublicReadPermissionWhereAuthor(ctx context.Context, id string, author string) error {
// 	query := `
// 		DELETE FROM reference
// 		WHERE
// 			from_id = $1
// 			AND to_id = '00000000-0000-0000-0000-000000000000'
// 			AND category = 'read'
// 			AND EXISTS (
// 				SELECT 1
// 				FROM knowledge
// 				WHERE
// 					id = $1
// 					AND author_id = $2
// 			)
// 	`

// 	_, err := single.DB.Exec(ctx, query, id, author)
// 	return err
// }
