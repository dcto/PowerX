package tag

import (
	"PowerX/internal/model/scrm/wechat/wecom/tag"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComTagGroupOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标签组列表/option
func NewListWeComTagGroupOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComTagGroupOptionLogic {
	return &ListWeComTagGroupOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComTagGroupOptionLogic) ListWeComTagGroupOption() (resp *types.ListWeComTagGroupReply, err error) {
	reply, _ := l.svcCtx.PowerX.SCRM.WeCom.FindListWeComTagGroupOption()

	return &types.ListWeComTagGroupReply{
		List: l.DTO(reply),
	}, err

}

// DTO
//
//	@Description:
//	@receiver group
//	@param groups
//	@return obj
func (l *ListWeComTagGroupOptionLogic) DTO(groups []*tag.WeComTagGroup) (obj []*types.WeComTagGroup) {

	if groups != nil {
		for _, g := range groups {
			obj = append(obj, &types.WeComTagGroup{
				GroupId:   g.GroupId,
				GroupName: g.Name,
			})
		}
	}
	return obj

}
