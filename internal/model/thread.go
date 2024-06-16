package model

import (
	"context"
	"encoding/json"
	"quanta/internal/utils"
)

// TODO: Whole thread model needs reworking

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	// TODO: this needs to be an edge
	Citations []string `json:"citations"`
}

type Thread []Message

func AppendMessage(ctx context.Context, threadID string, userID string, msg Message) error {
	_, err := CreateNodeWhereAuthor(ctx, msg.toNode(threadID, userID))
	return err
}

func UpdateMessage(ctx context.Context, id string, threadID string, userID string, msg Message) error {
	return UpdateNodeWhereAuthor(ctx, msg.toNode(threadID, userID))
}

func DeleteMessage(ctx context.Context, id string, userID string) error {
	return DeleteNodeWhereAuthor(ctx, id, userID)
}

func (m *Message) toNode(threadID string, userID string) Node {
	contentBytes, _ := json.Marshal(m)
	content := string(contentBytes)
	return Node{
		Parent:   threadID,
		Author:   userID,
		Category: "message",
		Title:    utils.EmptyString(),
		Content:  &content,
		Location: utils.EmptyString(),
	}
}
