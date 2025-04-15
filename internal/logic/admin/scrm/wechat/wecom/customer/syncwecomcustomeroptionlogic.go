package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncWeComCustomerOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量同步客户信息(根据员工ID同步/节流)
func NewSyncWeComCustomerOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWeComCustomerOptionLogic {
	return &SyncWeComCustomerOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWeComCustomerOptionLogic) SyncWeComCustomerOption(req *types.WeComCustomersRequest) (resp *types.WeComListCustomersReply, err error) {
	// todo: add your logic here and delete this line

	return
}
