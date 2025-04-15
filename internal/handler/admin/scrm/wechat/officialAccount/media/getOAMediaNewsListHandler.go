package media

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/officialAccount/media"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询菜单列表
func GetOAMediaNewsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := media.NewGetOAMediaNewsListLogic(r.Context(), svcCtx)
		resp, err := l.GetOAMediaNewsList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
