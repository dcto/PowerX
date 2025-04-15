package tag

import (
	tag2 "PowerX/internal/model/scrm/wechat/wecom/tag"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComTagGroupPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标签组分页/page
func NewListWeComTagGroupPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComTagGroupPageLogic {
	return &ListWeComTagGroupPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComTagGroupPageLogic) ListWeComTagGroupPage(req *types.ListWeComTagGroupPageRequest) (resp *types.ListWeComTagGroupPageReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.FindListWeComTagGroupPage(l.OPT(req))

	return &types.ListWeComTagGroupPageReply{
		List:      l.DTO(reply.List),
		PageIndex: reply.PageIndex,
		PageSize:  reply.PageSize,
		Total:     reply.Total,
	}, err

}

// @Description:
// @receiver group
// @param opt
// @return *types.PageOption[types.ListWeComTagGroupPageRequest]
func (l *ListWeComTagGroupPageLogic) OPT(opt *types.ListWeComTagGroupPageRequest) *types.PageOption[types.ListWeComTagGroupPageRequest] {

	option := types.PageOption[types.ListWeComTagGroupPageRequest]{
		Option: types.ListWeComTagGroupPageRequest{
			GroupId:   opt.GroupId,
			GroupName: opt.GroupName,
		},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	option.DefaultPageIfNotSet()

	return &option
}

// DTO
//
//	@Description:
//	@receiver group
//	@param datas
//	@return groups
func (l *ListWeComTagGroupPageLogic) DTO(datas []*tag2.WeComTagGroup) (groups []*types.GroupWithTag) {

	if datas == nil {
		return groups
	}
	for _, data := range datas {
		groups = append(groups, &types.GroupWithTag{
			GroupId:   data.GroupId,
			GroupName: data.Name,
			Tags:      l.tags(data.WeComGroupTags),
		})
	}
	return groups

}

// tags
//
//	@Description:
//	@receiver group
//	@param datas
//	@return tags
func (l *ListWeComTagGroupPageLogic) tags(datas []*tag2.WeComTag) (tags []*types.WeComTag) {

	if datas == nil {
		return tags
	}
	for _, data := range datas {
		tags = append(tags, &types.WeComTag{
			IsSelf: data.IsSelf,
			TagId:  data.TagId,
			Name:   data.Name,
			Sort:   data.Sort,
		})
	}
	return tags

}
