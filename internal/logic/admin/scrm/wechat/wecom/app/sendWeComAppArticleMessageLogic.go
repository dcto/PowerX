package app

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"

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

	option := l.OPT(req)

	_, err = l.svcCtx.PowerX.SCRM.WeCom.PushAppWeComMessageArticlesRequest(option, req.SendTime)

	return &types.AppMessageBaseReply{
		Message: `success`,
	}, err
}

// OPT
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *request.RequestMessageSendNews
func (l *SendWeComAppArticleMessageLogic) OPT(opt *types.AppMessageArticlesRequest) *request.RequestMessageSendNews {

	article := &request.RequestMessageSendNews{RequestMessageSend: request.RequestMessageSend{
		ToUser:  opt.ToUser,
		ToParty: opt.ToParty,
		ToTag:   opt.ToTag,
		MsgType: opt.MsgType,
		AgentID: opt.AgentID},
	}
	arc := &request.RequestNews{}
	for _, val := range opt.News.Article {
		arc.Article = append(arc.Article, &request.RequestNewsArticle{
			Title:       val.Title,
			Description: val.Description,
			URL:         val.URL,
			PicURL:      val.PicURL,
		})
	}
	article.News = arc
	return article
}
