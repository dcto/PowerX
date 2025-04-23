package tag

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncWeComGroupTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 全量同步标签/sync
func NewSyncWeComGroupTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWeComGroupTagLogic {
	return &SyncWeComGroupTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWeComGroupTagLogic) SyncWeComGroupTag() (resp *types.StatusWeComReply, err error) {
	_, err = l.svcCtx.PowerX.SCRM.WeCom.PullListWeComCorpTagRequest(nil, nil, 1)

	return &types.StatusWeComReply{
		Status: `success`,
	}, err

}
