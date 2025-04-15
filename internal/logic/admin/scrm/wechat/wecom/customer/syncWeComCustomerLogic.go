package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncWeComCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 同步客户
func NewSyncWeComCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWeComCustomerLogic {
	return &SyncWeComCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWeComCustomerLogic) SyncWeComCustomer() (resp *types.SyncWeComCustomerReply, err error) {
	// todo: add your logic here and delete this line

	return
}
