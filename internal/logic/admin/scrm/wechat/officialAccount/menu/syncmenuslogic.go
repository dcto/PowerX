package menu

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 请求菜单上传链接
func NewSyncMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMenusLogic {
	return &SyncMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncMenusLogic) SyncMenus(req *types.SyncMenusRequest) (resp *types.SyncMenusReply, err error) {
	// todo: add your logic here and delete this line

	return
}
