package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComCustomerPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 所有客户列表/page
func NewListWeComCustomerPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComCustomerPageLogic {
	return &ListWeComCustomerPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComCustomerPageLogic) ListWeComCustomerPage(req *types.WeComCustomersRequest) (resp *types.WeComListCustomersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
