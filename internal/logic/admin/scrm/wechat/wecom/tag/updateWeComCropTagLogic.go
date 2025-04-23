package tag

import (
	"context"
	tagReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWeComCropTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑企业标签
func NewUpdateWeComCropTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWeComCropTagLogic {
	return &UpdateWeComCropTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWeComCropTagLogic) UpdateWeComCropTag(req *types.UpdateCorpTagRequest) (resp *types.StatusWeComReply, err error) {
	cropTag, err := l.OPT(req)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.PowerX.SCRM.WeCom.UpdateWeComCorpTagRequest(cropTag)

	return &types.StatusWeComReply{
		Status: `success`,
	}, err
}

// (opt *types.UpdateCorpTagRequest)
//
//	@Description:
//	@receiver tag
//	@param opt
//	@return cropTag
//	@return err
func (l *UpdateWeComCropTagLogic) OPT(opt *types.UpdateCorpTagRequest) (cropTag *tagReq.RequestTagEditCorpTag, err error) {

	return &tagReq.RequestTagEditCorpTag{
		ID:      opt.TagId,
		Name:    opt.Name,
		Order:   opt.Sort,
		AgentID: &opt.AgentId,
	}, err
}
