package app

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/app"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// App企业群推送图文信息
func SendWeComAppGroupArticleMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppGroupMessageArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewSendWeComAppGroupArticleMessageLogic(r.Context(), svcCtx)
		resp, err := l.SendWeComAppGroupArticleMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
