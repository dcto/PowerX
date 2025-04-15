package bot

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BotWeComArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 机器人发送图文信息
func NewBotWeComArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BotWeComArticlesLogic {
	return &BotWeComArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BotWeComArticlesLogic) BotWeComArticles(req *types.GroupRobotMsgNewsArticlesRequest) (resp *types.GroupRobotMsgNewsArticlesReply, err error) {
	// todo: add your logic here and delete this line

	return
}
