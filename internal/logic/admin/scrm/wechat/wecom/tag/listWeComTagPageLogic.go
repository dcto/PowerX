package tag

import (
	"PowerX/internal/model/scrm/wechat/wecom/tag"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComTagPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标签列表/page
func NewListWeComTagPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComTagPageLogic {
	return &ListWeComTagPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComTagPageLogic) ListWeComTagPage(req *types.ListWeComTagReqeust) (resp *types.ListWeComTagReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.FindListWeComTagPage(l.OPT(req))
	if err != nil {
		return nil, err
	}

	return &types.ListWeComTagReply{
		List:      l.DTO(reply.List),
		PageIndex: reply.PageIndex,
		PageSize:  reply.PageSize,
		Total:     reply.Total,
	}, err

}

// @Description:
// @receiver tag
// @param opt
// @return *types.PageOption[types.ListWeComTagReqeust]
func (l *ListWeComTagPageLogic) OPT(opt *types.ListWeComTagReqeust) *types.PageOption[types.ListWeComTagReqeust] {

	option := types.PageOption[types.ListWeComTagReqeust]{
		Option:    types.ListWeComTagReqeust{},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	option.DefaultPageIfNotSet()
	if len(opt.TagIds) > 0 {
		option.Option.TagIds = opt.TagIds
	}
	if len(opt.GroupIds) > 0 {
		option.Option.GroupIds = opt.GroupIds
	}
	if opt.Name != `` {
		option.Option.Name = opt.Name
	}
	return &option

}

// DTO
//
//	@Description:
//	@receiver tag
//	@param tags
//	@return obj
func (l *ListWeComTagPageLogic) DTO(tags []*tag.WeComTag) (obj []*types.WeComTag) {

	if tags != nil {
		for _, val := range tags {
			obj = append(obj, &types.WeComTag{
				Type:      val.Type,
				IsSelf:    val.IsSelf,
				TagId:     val.TagId,
				GroupId:   val.GroupId,
				GroupName: val.WeComGroup.Name,
				Name:      val.Name,
				Sort:      val.Sort,
			})
		}
	}
	return obj

}
