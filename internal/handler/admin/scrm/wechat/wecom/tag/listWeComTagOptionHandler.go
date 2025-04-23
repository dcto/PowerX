package tag

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/tag"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 标签列表对象/key=>val
func ListWeComTagOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewListWeComTagOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListWeComTagOption()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
