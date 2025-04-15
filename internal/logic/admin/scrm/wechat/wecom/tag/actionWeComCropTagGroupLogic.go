package tag

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionWeComCropTagGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加、删除标签组内的标签
func NewActionWeComCropTagGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionWeComCropTagGroupLogic {
	return &ActionWeComCropTagGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionWeComCropTagGroupLogic) ActionWeComCropTagGroup(req *types.ActionCorpTagGroupRequest) (resp *types.StatusWeComReply, err error) {
	/*if len(opt.Tags) == 0 {
		return nil, errorx.ErrBadRequest
	}*/

	_, err = l.svcCtx.PowerX.SCRM.WeCom.ActionWeComCorpTagGroupRequest(req)
	if err != nil {
		return nil, err
	}

	return &types.StatusWeComReply{
		Status: `success`,
	}, err
}
