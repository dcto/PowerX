package organization

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/organization"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 部门列表/page
func ListWeComDepartmentPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListWeComDepartmentPageReqeust
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := organization.NewListWeComDepartmentPageLogic(r.Context(), svcCtx)
		resp, err := l.ListWeComDepartmentPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
