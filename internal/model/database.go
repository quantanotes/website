package model

// import (
// 	"context"
// 	"fmt"
// 	"quanta/internal/globals"
// 	"quanta/internal/services"
// 	"quanta/internal/utils"
// )

// type columnType string

// var (
// 	columnTypeBool columnType = "bool"
// 	columnTypeInt  columnType = "int"
// 	columnTypeNote columnType = "note"
// 	columnTypes               = map[columnType]bool{
// 		columnTypeBool: true,
// 		columnTypeInt:  true,
// 		columnTypeNote: true,
// 	}
// )

// var (
// 	metadataColumnKey = "column"
// )

// type Column struct {
// 	Name string
// 	Type string
// }

// func AddColumnWhereAuthor(ctx context.Context, nodeID string, authorID string, columnName string, columnType string) error {
// 	node, err := GetNodeWhereAuthor(ctx, nodeID, authorID)
// 	if err != nil {
// 		return err
// 	}
// 	utils.ReplaceNilWithEmptyMap(&node.Metadata)

// 	columns, ok := node.Metadata[metadataColumnKey].(map[string]any)
// 	if !ok {
// 		columns = make(map[string]any, 0)
// 	}

// 	if _, ok := columns[columnName]; ok {
// 		return globals.ErrBadRequest(fmt.Sprintf("Column %s already exists", columnName))
// 	}
// 	columns[columnName] = columnType
// 	node.Metadata[metadataColumnKey] = columns

// 	query := `
// 		UPDATE nodes
// 		SET metadata = $3
// 		WHERE id = $1 AND author_id = $2
// 	`
// 	_, err = services.DB.Exec(ctx, query, nodeID, authorID, node.Metadata)
// 	return err
// }

// func RemoveColumnWhereAuthor(ctx context.Context, nodeID string, authorID string, columnName string) error {
// 	node, err := GetNodeWhereAuthor(ctx, nodeID, authorID)
// 	if err != nil {
// 		return err
// 	}
// 	utils.ReplaceNilWithEmptyMap(&node.Metadata)

// 	columns, ok := node.Metadata[metadataColumnKey].(map[string]any)
// 	if !ok {
// 		return globals.ErrBadRequest(fmt.Sprintf("Column %s does not exist", columnName))
// 	}

// 	if _, ok := columns[columnName]; !ok {
// 		return globals.ErrBadRequest(fmt.Sprintf("Column %s does not exist", columnName))
// 	}
// 	delete(columns, columnName)
// 	node.Metadata[metadataColumnKey] = columns

// 	query := `
// 		UPDATE nodes
// 		SET metadata = $3
// 		WHERE id = $1 AND author_id = $2
// 	`
// 	_, err = services.DB.Exec(ctx, query, nodeID, authorID, node.Metadata)
// 	return err
// }

// func validColumnType(typ string) bool {
// 	_, ok := columnTypes[columnType(typ)]
// 	return ok
// }
