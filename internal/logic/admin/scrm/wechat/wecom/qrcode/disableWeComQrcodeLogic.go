package qrcode

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisableWeComQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 禁用场景码
func NewDisableWeComQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisableWeComQRCodeLogic {
	return &DisableWeComQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisableWeComQRCodeLogic) DisableWeComQRCode(req *types.ActionRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	if req.Qid == `` {
		return nil, errorx.ErrBadRequest
	}

	err = l.svcCtx.PowerX.SCRM.WeCom.ActionCustomerGroupQRCode(req.Qid, 2)
	if err != nil {
		return nil, errorx.ErrDeleteObject
	}

	return &types.ActionWeComGroupQRCodeActiveReply{
		Status: `success`,
	}, err
}
