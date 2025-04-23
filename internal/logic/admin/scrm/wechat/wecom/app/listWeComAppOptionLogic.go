package app

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/response"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComAppOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App列表/options
func NewListWeComAppOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComAppOptionLogic {
	return &ListWeComAppOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComAppOptionLogic) ListWeComAppOption() (resp *types.AppWeComListReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.PullListWeComAppRequest()

	return &types.AppWeComListReply{
		List: l.DTO(reply),
	}, err
}

// DTO
//
//	@Description:
//	@receiver this
//	@param list
//	@return apps
func (l *ListWeComAppOptionLogic) DTO(list *response.ResponseAgentList) (apps []*types.AppWechat) {
	for _, obj := range list.AgentList {
		apps = append(apps, &types.AppWechat{
			Agentid:       obj.AgentID,
			Name:          obj.Name,
			SquareLogoUrl: obj.SquareLogoURL,
		})
	}

	return apps
}
