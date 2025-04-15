package customer

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/customer"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 批量同步客户信息(根据员工ID同步/节流)
func SyncWeComCustomerOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WeComCustomersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewSyncWeComCustomerOptionLogic(r.Context(), svcCtx)
		resp, err := l.SyncWeComCustomerOption(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
