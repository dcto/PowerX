package brainx

import (
	"context"
	"net/http"
)

func (sp *BrainXServiceProvider) Chat(ctx context.Context, jsonBody interface{}, w http.ResponseWriter) error {
	url := "/chat-bot/chat"

	err := sp.Client.StreamPOST(ctx, url, jsonBody, w, nil)

	return err
}

func (sp *BrainXServiceProvider) AgentChat(ctx context.Context, jsonBody interface{}, w http.ResponseWriter) error {
	url := "/chat-bot/agent/chat"

	err := sp.Client.StreamPOST(ctx, url, jsonBody, w, nil)

	return err
}
