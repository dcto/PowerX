package qrcode

import (
	"PowerX/internal/types/errorx"
	"context"
	"fmt"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWeComQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新场景码
func NewUpdateWeComQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWeComQRCodeLogic {
	return &UpdateWeComQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWeComQRCodeLogic) UpdateWeComQRCode(req *types.QRCodeActiveRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	if err = l.OPT(req); err != nil {
		return nil, err
	}

	err = l.svcCtx.PowerX.SCRM.WeCom.UpdateWeComCustomerGroupQRCodeRequest(req)
	if err != nil {
		return nil, errorx.ErrCreateObject
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
//	@return err
func (l *UpdateWeComQRCodeLogic) OPT(opt *types.QRCodeActiveRequest) (err error) {

	if opt.Name == `` {
		err = fmt.Errorf(`Name error`)
	} else if opt.SceneLink == `` {
		err = fmt.Errorf(`SceneLink error`)
	} else if opt.RealQRCodeLink == `` {
		err = fmt.Errorf(`RealQRCode error`)
	} else if opt.ExpiryDate == 0 {
		err = fmt.Errorf(`ExpiryDate error`)
	} else if opt.Owner == nil {
		err = fmt.Errorf(`Owner error`)
	} else if opt.Qid == `` {
		err = fmt.Errorf(`Qid error`)
	}

	return err
}
