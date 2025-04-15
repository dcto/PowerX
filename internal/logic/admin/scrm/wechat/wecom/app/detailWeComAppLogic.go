package app

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/response"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailWeComAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App详情
func NewDetailWeComAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailWeComAppLogic {
	return &DetailWeComAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailWeComAppLogic) DetailWeComApp(req *types.ApplicationRequest) (resp *types.ApplicationReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.PullDetailWeComAppRequest(req.AgentId)

	return l.DTO(reply), err
}

// DTO
//
//	@Description:
//	@receiver this
//	@param detail
//	@return *types.ApplicationReply
func (l *DetailWeComAppLogic) DTO(detail *response.ResponseAgentGet) *types.ApplicationReply {

	return &types.ApplicationReply{
		Agentid:                 detail.AgentID,
		Name:                    detail.Name,
		SquareLogoUrl:           detail.SquareLogoURL,
		Description:             detail.Description,
		AllowUserinfos:          l.allowUserInfos(detail.AllowUserInfos),
		AllowPartys:             l.allowPartys(detail.AllowParty),
		AllowTags:               l.allowTags(detail.AllowTags),
		Close:                   int(detail.Close),
		RedirectDomain:          detail.RedirectDomain,
		ReportLocationFlag:      int(detail.ReportLocationFlag),
		Isreportenter:           int(detail.IsReportEnter),
		HomeUrl:                 detail.HomeURL,
		CustomizedPublishStatus: 0,
	}

}

// allowUserInfos
//
//	@Description:
//	@receiver this
func (l *DetailWeComAppLogic) allowUserInfos(infos response.ResponseAgentAllowUserInfos) (infox types.AllowUserinfos) {

	for _, user := range infos.User {
		infox.User = append(infox.User, types.WeComUser{
			Userid: user.UserID,
		})
	}
	return infox

}

// allowPartys
//
//	@Description:
//	@receiver this
//	@param party
//	@return types.AllowPartys
func (l *DetailWeComAppLogic) allowPartys(party response.ResponseAgentAllowParty) types.AllowPartys {
	return types.AllowPartys{
		Partyid: party.PartyID,
	}
}

// allowTags
//
//	@Description:
//	@receiver this
//	@param tags
//	@return types.AllowTags
func (l *DetailWeComAppLogic) allowTags(tags response.ResponseAgentAllowTags) types.AllowTags {
	return types.AllowTags{
		Tagid: tags.TagID,
	}
}
