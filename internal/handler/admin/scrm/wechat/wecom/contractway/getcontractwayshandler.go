package contractway

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/contractway"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询渠道活码
func GetContractWaysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetContractWaysRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contractway.NewGetContractWaysLogic(r.Context(), svcCtx)
		resp, err := l.GetContractWays(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
