package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComCustomerGroupLimitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户群列表/limit
func NewListWeComCustomerGroupLimitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComCustomerGroupLimitLogic {
	return &ListWeComCustomerGroupLimitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComCustomerGroupLimitLogic) ListWeComCustomerGroupLimit(req *types.WeComCustomerGroupRequest) (resp *types.WeComListCustomerGroupReply, err error) {
	// todo: add your logic here and delete this line

	return
}
