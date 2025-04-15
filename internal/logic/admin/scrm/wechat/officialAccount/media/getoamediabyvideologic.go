package media

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOAMediaByVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据媒体key获取媒体
func NewGetOAMediaByVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAMediaByVideoLogic {
	return &GetOAMediaByVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOAMediaByVideoLogic) GetOAMediaByVideo(req *types.GetOAMediaRequest) (resp *types.GetOAMediaByVideoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
