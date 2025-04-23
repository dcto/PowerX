package resource

import (
	"PowerX/internal/model/scrm/wechat/wecom/resource"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信素材库/page
func NewListWeComImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComImageLogic {
	return &ListWeComImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComImageLogic) ListWeComImage(req *types.ListWeComResourceImageRequest) (resp *types.ListWeComResourceImageReply, err error) {
	page, err := l.svcCtx.PowerX.SCRM.WeCom.FindWeComResourceListFromLocalPage(req)

	return &types.ListWeComResourceImageReply{
		List:      l.DTO(page.List),
		PageIndex: page.PageIndex,
		PageSize:  page.PageSize,
		Total:     page.Total,
	}, err
}

// DTO
//
//	@Description:
//	@receiver image
//	@param data
//	@return resources
func (l *ListWeComImageLogic) DTO(data []*resource.WeComResource) (resources []*types.Resource) {

	if data != nil {
		for _, obj := range data {
			resources = append(resources, &types.Resource{
				Link:         obj.Url,
				ResourceType: obj.ResourceType,
				CreateTime:   obj.CreatedAt.String(),
			})
		}
	}
	return resources

}
