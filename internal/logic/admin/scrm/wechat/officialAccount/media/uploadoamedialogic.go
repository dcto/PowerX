package media

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadOAMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建菜单
func NewUploadOAMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadOAMediaLogic {
	return &UploadOAMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadOAMediaLogic) UploadOAMedia() (resp *types.CreateOAMediaReply, err error) {
	// todo: add your logic here and delete this line

	return
}
