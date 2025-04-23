package wecom

import (
	wecomLogic "PowerX/internal/logic/webhook/wecom"
	"net/http"

	"PowerX/internal/svc"
)

func PostMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := wecomLogic.NewWebhookPostMessageLogic(r.Context(), svcCtx)
		l.WebhookPostMessage(w, r)

	}
}
