package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComAppGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App企业群列表/list
func NewListWeComAppGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComAppGroupLogic {
	return &ListWeComAppGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComAppGroupLogic) ListWeComAppGroup(req *types.AppGroupListRequest) (resp *types.AppGroupListReply, err error) {
	var chatIds []string
	if req.ChatId != `` {
		chatIds = append(chatIds, req.ChatId)
	}
	replies, err := l.svcCtx.PowerX.SCRM.WeCom.PullListWeComAppGroupRequest(chatIds...)

	return &types.AppGroupListReply{
		replies,
	}, err
}
