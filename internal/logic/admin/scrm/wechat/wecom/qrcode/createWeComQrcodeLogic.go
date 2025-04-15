package qrcode

import (
	"PowerX/internal/types/errorx"
	"PowerX/pkg/idx"
	"context"
	"fmt"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWeComQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建场景码
func NewCreateWeComQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWeComQRCodeLogic {
	return &CreateWeComQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWeComQRCodeLogic) CreateWeComQRCode(req *types.QRCodeActiveRequest) (resp *types.ActionWeComGroupQRCodeActiveReply, err error) {
	if err = l.OPT(req); err != nil {
		return nil, err
	}

	err = l.svcCtx.PowerX.SCRM.WeCom.CreateWeComCustomerGroupQRCodeRequest(req)
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
func (l *CreateWeComQRCodeLogic) OPT(opt *types.QRCodeActiveRequest) (err error) {

	generate, err := idx.Generate()
	if err != nil {
		err = fmt.Errorf(`Qid error`)
	} else {
		opt.Qid = generate
	}

	return err
}
