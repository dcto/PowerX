package wecom

import (
	wecomLogic "PowerX/internal/logic/webhook/wecom"
	"net/http"

	"PowerX/internal/svc"
)

func GetMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := wecomLogic.NewWebhookGetMessageLogic(r.Context(), svcCtx)
		l.WebhookGetMessage(w, r)

	}
}
