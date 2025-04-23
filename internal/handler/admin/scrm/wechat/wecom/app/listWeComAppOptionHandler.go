package app

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/app"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// App列表/options
func ListWeComAppOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := app.NewListWeComAppOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListWeComAppOption()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
