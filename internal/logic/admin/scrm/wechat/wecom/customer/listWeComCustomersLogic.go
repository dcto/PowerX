package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComCustomersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户详情列表
func NewListWeComCustomersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComCustomersLogic {
	return &ListWeComCustomersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComCustomersLogic) ListWeComCustomers(req *types.ListWeComCustomersRequest) (resp *types.ListWeComCustomersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
