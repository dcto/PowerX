package resource

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/resource"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传图片到微信
func CreateWeComImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := resource.NewCreateWeComImageLogic(r.Context(), svcCtx)
		resp, err := l.CreateWeComImage(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
