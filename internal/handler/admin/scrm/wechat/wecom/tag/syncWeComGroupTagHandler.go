package tag

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/tag"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 全量同步标签/sync
func SyncWeComGroupTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewSyncWeComGroupTagLogic(r.Context(), svcCtx)
		resp, err := l.SyncWeComGroupTag()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
