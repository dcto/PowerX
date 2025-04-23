package organization

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取员工列表
func NewListWeComUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComUserLogic {
	return &ListWeComUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComUserLogic) ListWeComUser(req *types.ListWeComUserReqeust) (resp *types.ListWeComUserReply, err error) {
	// todo: add your logic here and delete this line

	return
}
