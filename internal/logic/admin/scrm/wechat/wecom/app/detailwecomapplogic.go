package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailWeComAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App详情
func NewDetailWeComAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailWeComAppLogic {
	return &DetailWeComAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailWeComAppLogic) DetailWeComApp(req *types.ApplicationRequest) (resp *types.ApplicationReply, err error) {
	// todo: add your logic here and delete this line

	return
}
