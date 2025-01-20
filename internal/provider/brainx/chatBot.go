package brainx

import (
	"PowerX/internal/logic/openapi/provider/brainx/schema"
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"net/http"
)

func DefaultCommandCallback(message *schema.SSEMessage) (*schema.SSEMessage, error) {
	if message.Status == schema.StatusFinished {
		message.Command = schema.Command{
			Type: "render",
			Payload: object.HashMap{
				"arg1": 1,
				"arg2": "run",
			},
		}
	}
	return message, nil
}

func (sp *BrainXServiceProvider) Chat(ctx context.Context, jsonBody interface{}, w http.ResponseWriter) error {
	url := "/chat-bot/chat"
	commandFun := DefaultCommandCallback

	err := sp.Client.StreamPOST(ctx, url, jsonBody, w, nil, commandFun)

	return err
}

func (sp *BrainXServiceProvider) AgentChat(ctx context.Context, jsonBody interface{}, w http.ResponseWriter) error {
	url := "/chat-bot/agent/chat"

	commandFun := DefaultCommandCallback

	err := sp.Client.StreamPOST(ctx, url, jsonBody, w, nil, commandFun)

	return err
}
