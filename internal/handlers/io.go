package handlers

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
)

type ioChatRequest struct {
	Thread   []model.Message
	ThreadID string `json:"threadID"`
}

var ioChat = protectedJSONRequestWithWriterHandler(func(ctx context.Context, w http.ResponseWriter, userID string, req ioChatRequest) error {
	ioRes, err := callIo(req)
	if err != nil {
		return err
	}
	defer ioRes.Close()

	ioMsg, err := writeIoChat(w, ioRes)
	if err != nil {
		return err
	}

	userMsg := req.Thread[len(req.Thread)-1]
	if err := model.AppendMessage(ctx, req.ThreadID, userID, userMsg); err != nil {
		return err
	}

	if err := model.AppendMessage(ctx, req.ThreadID, userID, ioMsg); err != nil {
		return err
	}

	return nil
})

func callIo(req ioChatRequest) (io.ReadCloser, error) {
	threadJson, _ := json.Marshal(req.Thread)
	threadReader := bytes.NewReader(threadJson)

	res, err := http.Post(globals.IoURL+"/chat", "application/json", threadReader)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		var errMsg []byte
		threadReader.Read(errMsg)
		return nil, errors.New(string(errMsg))
	}

	return res.Body, nil
}

func writeIoChat(w io.Writer, body io.ReadCloser) (model.Message, error) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return model.Message{}, globals.NewError(http.StatusHTTPVersionNotSupported, "Streaming not supported")
	}

	var msg model.Message
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		w.Write([]byte(line + "\n"))

		var part model.Message
		if err := json.Unmarshal([]byte(line), &part); err != nil {
			continue
		}

		msg.Content += part.Content
		msg.Role = "assistant" // Temporary until I can figure out where role is being set to system incorrectly
		msg.Citations = append(msg.Citations, part.Citations...)

		flusher.Flush()
	}

	return msg, nil
}
