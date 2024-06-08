package model

import (
	"context"
	"quanta/internal/single"
	"time"
)

type Knowledge struct {
	Id         string    `json:"id"`
	Parent     string    `json:"parent"`
	Author     string    `json:"author"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Title      *string   `json:"title"`
	Content    *string   `json:"content"`
	Permission int       `json:"permission"`
}

func GetKnowledge(ctx context.Context, id string) (Knowledge, error) {
	var k Knowledge
	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8'), permission,
		FROM knowledge
		WHERE id = $1 AND category = 'note'
	`
	err := single.DB.QueryRow(ctx, query, id).Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content, &k.Permission)
	return k, err
}

func CreateKnowledge(ctx context.Context, parent string, author string, category string, title string, content string) (Knowledge, error) {
	var k Knowledge
	query := `
		INSERT INTO knowledge (parent_id, author_id, category, title, content)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, parent_id, author_id, title, convert_from(content, 'utf-8')
	`
	err := single.DB.
		QueryRow(ctx, query, parent, author, category, title, []byte(content)).
		Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content)
	return k, err
}

func DeleteKnowledgeWhereAuthor(ctx context.Context, id string, author string) error {
	query := `
		DELETE FROM knowledge
		WHERE id = $1 AND author_id = $2
	`
	_, err := single.DB.Exec(ctx, query, id, author)
	return err
}

func UpdateKnowledgeWhereAuthor(ctx context.Context, id string, author string, title string, content string) error {
	query := `
		UPDATE knowledge
		SET title = $3, content = convert_to($4, 'utf8')
		WHERE id = $1 AND author_id = $2	
	`
	_, err := single.DB.Exec(ctx, query, id, author, title, content)
	return err
}

func MoveKnowledgeWhereAuthor(ctx context.Context, id string, author string, to string) error {
	query := `
		UPDATE knowledge k
		SET parent_id = $3
		WHERE id = $1 AND author_id = $2
			AND EXISTS (
				SELECT * FROM knowledge WHERE id = k.parent_id AND author_id = $2
			)
	`

	_, err := single.DB.Exec(ctx, query, id, author, to)
	return err
}

func CreateKnowledgeWhereParentAuthor(
	ctx context.Context, parent string, author string, category string, title string, content string, permission int,
) (Knowledge, error) {
	var k Knowledge
	query := `
		INSERT INTO knowledge (parent_id, author_id, category, title, content, permission)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, parent_id, author_id, title, convert_from(content, 'utf-8'), permission 
	`
	err := single.DB.
		QueryRow(ctx, query, parent, author, category, title, []byte(content), permission).
		Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content, &k.Permission)
	return k, err
}

func GetKnowledgeWherePublic(ctx context.Context) ([]Knowledge, error) {
	var ks []Knowledge

	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8')
		FROM knowledge
		WHERE category = 'note' AND permission > 0
	`

	rows, err := single.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k Knowledge
		if err := rows.Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content); err != nil {
			return nil, err
		}
		ks = append(ks, k)
	}

	return ks, rows.Err()
}

func GetKnowledgeChildren(ctx context.Context, id string) ([]Knowledge, error) {
	var ks []Knowledge

	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8')
		FROM knowledge
		WHERE parent_id = $1 AND category = 'note'
	`

	rows, err := single.DB.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k Knowledge
		if err := rows.Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content); err != nil {
			return nil, err
		}
		ks = append(ks, k)
	}

	return ks, rows.Err()
}

func GetKnowledgeChildWhereAuthorCategory(ctx context.Context, parent string, author string, category string) (Knowledge, error) {
	var k Knowledge
	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8'), permission 
		FROM knowledge
		WHERE
			parent_id = $1 AND
			author_id = $2 AND
			category = $3
	`
	err := single.DB.
		QueryRow(ctx, query, parent, author, category).
		Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content, &k.Permission)
	return k, err
}

func GetKnowledgeChildrenWhereAuthor(ctx context.Context, parent string, author string) ([]Knowledge, error) {
	var ks []Knowledge

	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8'), permission
		FROM knowledge
		WHERE
			parent_id = $1 AND
			author_id = $2
	`

	rows, err := single.DB.Query(ctx, query, parent, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k Knowledge
		if err := rows.Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content, &k.Permission); err != nil {
			return nil, err
		}
		ks = append(ks, k)
	}

	return ks, rows.Err()
}

func GetKnowledgeChildrenWhereAuthorCategory(ctx context.Context, parent string, author string, category string) ([]Knowledge, error) {
	var ks []Knowledge

	query := `
		SELECT id, parent_id, author_id, title, convert_from(content, 'utf-8'), permission 
		FROM knowledge
		WHERE
			parent_id = $1 AND
			author_id = $2 AND
			category = $3
	`

	rows, err := single.DB.Query(ctx, query, parent, author, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k Knowledge
		if err := rows.Scan(&k.Id, &k.Parent, &k.Author, &k.Title, &k.Content, &k.Permission); err != nil {
			return nil, err
		}
		ks = append(ks, k)
	}

	return ks, rows.Err()
}
