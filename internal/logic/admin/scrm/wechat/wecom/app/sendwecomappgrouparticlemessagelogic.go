package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendWeComAppGroupArticleMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App企业群推送图文信息
func NewSendWeComAppGroupArticleMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendWeComAppGroupArticleMessageLogic {
	return &SendWeComAppGroupArticleMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendWeComAppGroupArticleMessageLogic) SendWeComAppGroupArticleMessage(req *types.AppGroupMessageArticleRequest) (resp *types.AppGroupMessageReply, err error) {
	// todo: add your logic here and delete this line

	return
}
