package tag

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/tag"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 标签组列表/option
func ListWeComTagGroupOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewListWeComTagGroupOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListWeComTagGroupOption()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
