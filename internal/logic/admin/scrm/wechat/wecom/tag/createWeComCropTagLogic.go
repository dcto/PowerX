package tag

import (
	"context"
	tagReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWeComCropTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建企业标签
func NewCreateWeComCropTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWeComCropTagLogic {
	return &CreateWeComCropTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWeComCropTagLogic) CreateWeComCropTag(req *types.CreateCorpTagRequest) (resp *types.StatusWeComReply, err error) {
	cropTag, err := l.OPT(req)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.PowerX.SCRM.WeCom.CreateWeComCorpTagRequest(cropTag)

	return &types.StatusWeComReply{
		Status: `success`,
	}, err

}

// (opt *types.CreateCorpTagRequest)
//
//	@Description:
//	@receiver tag
//	@param opt
//	@return cropTag
//	@return err
func (l *CreateWeComCropTagLogic) OPT(opt *types.CreateCorpTagRequest) (cropTag *tagReq.RequestTagAddCorpTag, err error) {

	cropTag = &tagReq.RequestTagAddCorpTag{
		GroupID:   &opt.GroupId,
		GroupName: opt.GroupName,
		Order:     opt.Sort,
		Tag:       l.loadTagField(opt.Tag),
		AgentID:   &opt.AgentId,
	}
	if opt.GroupId == `` {
		cropTag.GroupID = nil
	}
	return cropTag, err
}

// loadTagFeild
//
//	@Description:
//	@receiver tag
//	@param tags
//	@return obj
func (l *CreateWeComCropTagLogic) loadTagField(tags []*types.TagFieldTag) (obj []tagReq.RequestTagAddCorpTagFieldTag) {
	for _, val := range tags {
		obj = append(obj, tagReq.RequestTagAddCorpTagFieldTag{
			Name:  val.Name,
			Order: val.Sort,
		})
	}
	return obj
}
