package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeComCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户详情
func NewGetWeComCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeComCustomerLogic {
	return &GetWeComCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWeComCustomerLogic) GetWeComCustomer(req *types.GetWeComCustomerRequest) (resp *types.GetWeComCustomerReply, err error) {
	// todo: add your logic here and delete this line

	return
}
