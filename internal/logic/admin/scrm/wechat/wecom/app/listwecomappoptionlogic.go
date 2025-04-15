package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComAppOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App列表/options
func NewListWeComAppOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComAppOptionLogic {
	return &ListWeComAppOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComAppOptionLogic) ListWeComAppOption() (resp *types.AppWeComListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
