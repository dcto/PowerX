package qrcode

import (
	"PowerX/internal/types/errorx"
	"context"
	"fmt"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateActiveQRCodeLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下载场景码/upload
func NewUpdateActiveQRCodeLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateActiveQRCodeLinkLogic {
	return &UpdateActiveQRCodeLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateActiveQRCodeLinkLogic) UpdateActiveQRCodeLink(req *types.ActionRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	err = l.OPT(req)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.PowerX.SCRM.WeCom.UpdateSceneQRCodeLink(req.Qid, req.SceneQRCodeLink)
	if err != nil {
		return nil, errorx.ErrBadRequest
	}

	return &types.ActionWeComGroupQRCodeActiveReply{
		Status: `success`,
	}, err
}

// OPT
//
//	@Description:
//	@receiver qrcode
//	@param opt
//	@return error
func (l *UpdateActiveQRCodeLinkLogic) OPT(opt *types.ActionRequest) error {

	if opt.Qid == `` {
		return fmt.Errorf(`Qid error`)
	}
	if opt.SceneQRCodeLink == `` {
		return fmt.Errorf(`SceneQRCodeLink error`)
	}
	return nil

}
