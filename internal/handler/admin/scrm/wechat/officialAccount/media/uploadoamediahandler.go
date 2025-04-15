package media

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/officialAccount/media"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建菜单
func UploadOAMediaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := media.NewUploadOAMediaLogic(r.Context(), svcCtx)
		resp, err := l.UploadOAMedia()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
