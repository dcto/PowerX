package organization

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncWeComUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 同步组织架构/department&user
func NewSyncWeComUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWeComUserLogic {
	return &SyncWeComUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWeComUserLogic) SyncWeComUser() (resp *types.SyncWeComOrganizationReply, err error) {

	err = l.svcCtx.PowerX.SCRM.WeCom.PullSyncDepartmentsAndUsersRequest(l.ctx)

	return &types.SyncWeComOrganizationReply{
		Status: `success`,
	}, err
}
