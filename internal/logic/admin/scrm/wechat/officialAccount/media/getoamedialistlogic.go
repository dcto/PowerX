package media

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOAMediaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询菜单列表
func NewGetOAMediaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAMediaListLogic {
	return &GetOAMediaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOAMediaListLogic) GetOAMediaList(req *types.GetOAMediaListRequest) (resp *types.GetOAMediaListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
