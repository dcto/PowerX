package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendWeComAppArticleMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App发送图文信息
func NewSendWeComAppArticleMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendWeComAppArticleMessageLogic {
	return &SendWeComAppArticleMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendWeComAppArticleMessageLogic) SendWeComAppArticleMessage(req *types.AppMessageArticlesRequest) (resp *types.AppMessageBaseReply, err error) {
	// todo: add your logic here and delete this line

	return
}
