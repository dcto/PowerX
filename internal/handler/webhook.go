package handler

import (
	"PowerX/internal/handler/webhook/payment"
	"PowerX/internal/handler/webhook/wecom"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterWebhookHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/message",
					Handler: wecom.GetMessageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/message",
					Handler: wecom.PostMessageHandler(serverCtx),
				},
			}...,
		),
		//fixme /api/webhook/wecom,  Reverse Proxy api. [Eros]
		//rest.WithPrefix("/webhook/wecom"),
		rest.WithPrefix("/api/webhook/wecom"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/pay/",
					Handler: payment.PostMessageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/webhook/wx"),
	)

	// custom
}
