package qrcode

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWeComQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除场景码
func NewDeleteWeComQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWeComQRCodeLogic {
	return &DeleteWeComQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWeComQRCodeLogic) DeleteWeComQRCode(req *types.ActionRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	if req.Qid == `` {
		return nil, errorx.ErrBadRequest
	}

	err = l.svcCtx.PowerX.SCRM.WeCom.ActionCustomerGroupQRCode(req.Qid, 3)
	if err != nil {
		return nil, errorx.ErrDeleteObject
	}

	return &types.ActionWeComGroupQRCodeActiveReply{
		Status: `success`,
	}, err
}
