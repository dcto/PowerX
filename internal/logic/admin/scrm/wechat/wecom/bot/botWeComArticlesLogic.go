package bot

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/request"

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
	articles := []*request.GroupRobotMsgNewsArticles{
		{Title: req.Title, Description: req.Description, Url: req.Url, PicUrl: req.PicUrl},
	}
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.PushWeComBotArticlesRequest(req.Key, articles)
	if err != nil {
		return nil, err
	}
	resp.Messaage = reply.Message

	return resp, err
}
