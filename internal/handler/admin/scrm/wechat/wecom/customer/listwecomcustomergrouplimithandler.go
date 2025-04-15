package customer

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/customer"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 客户群列表/limit
func ListWeComCustomerGroupLimitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WeComCustomerGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewListWeComCustomerGroupLimitLogic(r.Context(), svcCtx)
		resp, err := l.ListWeComCustomerGroupLimit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
