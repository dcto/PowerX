package qrcode

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableWeComQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 启用场景码
func NewEnableWeComQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableWeComQRCodeLogic {
	return &EnableWeComQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnableWeComQRCodeLogic) EnableWeComQRCode(req *types.ActionRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	if req.Qid == `` {
		return nil, errorx.ErrBadRequest
	}

	err = l.svcCtx.PowerX.SCRM.WeCom.ActionCustomerGroupQRCode(req.Qid, 1)
	if err != nil {
		return nil, errorx.ErrDeleteObject
	}

	return &types.ActionWeComGroupQRCodeActiveReply{
		Status: `success`,
	}, err
}
