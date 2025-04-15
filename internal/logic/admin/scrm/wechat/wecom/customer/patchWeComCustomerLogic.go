package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchWeComCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改客户信息
func NewPatchWeComCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchWeComCustomerLogic {
	return &PatchWeComCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchWeComCustomerLogic) PatchWeComCustomer(req *types.PatchWeComCustomerRequest) (resp *types.PatchWeComCustomerReply, err error) {
	// todo: add your logic here and delete this line

	return
}
