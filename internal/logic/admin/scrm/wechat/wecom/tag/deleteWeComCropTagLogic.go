package tag

import (
	"PowerX/internal/types/errorx"
	"context"
	tagReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWeComCropTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除企业标签
func NewDeleteWeComCropTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWeComCropTagLogic {
	return &DeleteWeComCropTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWeComCropTagLogic) DeleteWeComCropTag(req *types.DeleteCorpTagRequest) (resp *types.StatusWeComReply, err error) {
	option, err := l.OPT(req)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.PowerX.SCRM.WeCom.DeleteWeComCorpTagRequest(option)

	return &types.StatusWeComReply{
		Status: `success`,
	}, err

}

// OPT
//
//	@Description:
//	@receiver tag
//	@param opt
//	@return *tagReq.RequestTagDelCorpTag
//	@return error
func (l *DeleteWeComCropTagLogic) OPT(opt *types.DeleteCorpTagRequest) (*tagReq.RequestTagDelCorpTag, error) {
	if opt == nil {
		return nil, errorx.ErrBadRequest
	}
	return &tagReq.RequestTagDelCorpTag{
		TagID:   opt.TagIds,
		GroupID: opt.GroupIds,
		AgentID: &opt.AgentId,
	}, nil
}
