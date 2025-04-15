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
	// todo: add your logic here and delete this line

	return
}
