package app

import (
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc/powerx/scrm/wechat/wecom"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

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
	option, err := l.OPT(req)
	if err != nil {
		return nil, errorx.ErrBadRequest
	}
	_, err = l.svcCtx.PowerX.SCRM.WeCom.PushAppWeComGroupMessageArticlesRequest(option, req.SendTime)

	return &types.AppGroupMessageReply{
		ChatIds: req.ChatIds,
	}, err

}

// OPT
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *power.HashMap
//	@return error
func (l *SendWeComAppGroupArticleMessageLogic) OPT(opt *types.AppGroupMessageArticleRequest) (*power.HashMap, error) {
	option := wecom.WechatAppRequestBase{
		ChatIds: opt.ChatIds,
		MsgType: `news`,
		Safe:    0,
		News:    wecom.WechatAppRequestNews{},
	}
	arc := wecom.WechatAppRequestNews{}
	arc.Article = append(arc.Article, &wecom.WechatAppRequestNewsArticle{
		Title:       opt.Title,
		Description: opt.Description,
		URL:         opt.URL,
		PicURL:      opt.PicURL,
	})
	option.News = arc
	hashMap, err := power.StructToHashMap(option)

	return hashMap, err
}
