package model

import (
	"context"
	"quanta/internal/services"
	"quanta/internal/utils"
	"time"
)

var (
	nodeReadColumns = `id, parent_id, author_id, created_at, updated_at, category, title, convert_from(content, 'utf-8'), location, permission`
)

type Node struct {
	ID         string    `json:"id" db:"id"`
	Parent     string    `json:"parent" db:"parent_id"`
	Author     string    `json:"author" db:"author_id"`
	Created    time.Time `json:"created" db:"created_at"`
	Updated    time.Time `json:"updated" db:"updated_at"`
	Category   string    `json:"category"`
	Title      *string   `json:"title"`
	Content    *string   `json:"content"`
	Location   *string   `json:"location"`
	Permission int       `json:"permission"`
}

func GetNodeWhereAuthor(ctx context.Context, id string, author string) (Node, error) {
	var node Node
	query := `
		SELECT ` + nodeReadColumns + ` 
		FROM nodes
		WHERE id = $1 AND author_id = $2
	`
	err := services.DB.QueryRow(ctx, query, id).Scan(node.scan()...)
	return node, err
}

func GetNodeChildrenWhereAuthor(ctx context.Context, parent string, author string) ([]Node, error) {
	query := `
		SELECT ` + nodeReadColumns + ` 
		FROM nodes
		WHERE parent_id = $1 AND author_id = $2
	`
	rows, err := services.DB.Query(ctx, query, parent, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res, err := scanMany[Node](rows)
	return res, err
}

func GetNodeChildWhereAuthorCategory(ctx context.Context, parent string, author string, category string) (Node, error) {
	var node Node
	query := `
		SELECT ` + nodeReadColumns + ` 
		FROM nodes
		WHERE parent_id = $1 AND author_id = $2 AND category = $3
	`
	err := services.
		DB.
		QueryRow(ctx, query, parent, author, category).
		Scan(node.scan()...)
	return node, err
}

func GetNodesWherePublic(ctx context.Context) ([]Node, error) {
	// Enforce category until we can update the space frontend
	query := `
		SELECT ` + nodeReadColumns + ` 
		FROM nodes
		WHERE category = 'note' AND permission > 0
	`
	rows, err := services.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res, err := scanMany[Node](rows)
	return res, err
}

func GetNodeChildrenWherePublic(ctx context.Context, parent string) ([]Node, error) {
	query := `
		SELECT ` + nodeReadColumns + ` 
		FROM nodes
		WHERE parent_id = $1 AND category = 'note' AND permission > 0
	`
	rows, err := services.DB.Query(ctx, query, parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res, err := scanMany[Node](rows)
	return res, err
}

func CreateNodeWhereAuthor(ctx context.Context, node Node) (Node, error) {
	query := `
		WITH check_parent AS (
			SELECT 1 FROM nodes WHERE id = $1 AND author_id = $2
		)
		INSERT INTO nodes (parent_id, author_id, category, title, content)
		SELECT $1, $2, $3, $4, convert_to($5, 'utf8')
		FROM check_parent
		RETURNING  
	` + nodeReadColumns
	err := services.
		DB.
		QueryRow(ctx, query, node.Parent, node.Author, node.Category, node.Title, []byte(*node.Content)).
		Scan(node.scan()...)
	return node, err
}

func UpdateNodeWhereAuthor(ctx context.Context, node Node) error {
	query := `
		UPDATE nodes
		SET title = $3, content = convert_to($4, 'utf8')
		WHERE id = $1 AND author_id = $2	
	`
	_, err := services.DB.Exec(ctx, query, node.ID, node.Author, node.Title, node.Content)
	return err
}

func DeleteNodeWhereAuthor(ctx context.Context, id string, author string) error {
	query := `
		DELETE FROM nodes
		WHERE id = $1 AND author_id = $2
	`
	_, err := services.DB.Exec(ctx, query, id, author)
	return err
}

func MoveNodeWhereAuthor(ctx context.Context, id string, author string, to string) error {
	query := `
		UPDATE nodes n
		SET parent_id = $3
		WHERE id = $1 AND author_id = $2 AND EXISTS (
			SELECT NULL FROM nodes WHERE id = n.parent_id AND author_id = $2
		)
	`
	_, err := services.DB.Exec(ctx, query, id, author, to)
	return err
}

func UpdateNodePermissionWhereAuthor(ctx context.Context, id string, author string, permission int) error {
	query := `
		UPDATE nodes
		SET permission = $3
		WHERE id = $1 AND author_id = $2
	`
	_, err := services.DB.Exec(ctx, query, id, author, permission)
	return err
}

func (n *Node) ReplaceNilWithEmpty() {
	utils.ReplaceNilWithEmptyString(&n.Title)
	utils.ReplaceNilWithEmptyString(&n.Location)
	utils.ReplaceNilWithEmptyString(&n.Content)
}

func (n *Node) scan() []any {
	n.ReplaceNilWithEmpty()
	return []any{
		&n.ID,
		&n.Parent,
		&n.Author,
		&n.Created,
		&n.Updated,
		&n.Category,
		&n.Title,
		&n.Content,
		&n.Location,
		&n.Permission,
	}
}
