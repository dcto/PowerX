package qrcode

import (
	"net/http"

	"PowerX/internal/logic/admin/scrm/wechat/wecom/qrcode"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 下载场景码/upload
func UpdateActiveQRCodeLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := qrcode.NewUpdateActiveQRCodeLinkLogic(r.Context(), svcCtx)
		resp, err := l.UpdateActiveQRCodeLink(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
