package organization

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/organization"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 同步组织架构/department&user
func SyncWeComUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := organization.NewSyncWeComUserLogic(r.Context(), svcCtx)
		resp, err := l.SyncWeComUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
