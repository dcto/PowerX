package tag

import (
	"PowerX/internal/types/errorx"
	"context"
	tagReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionWeComCustomerTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑/删除客户标签
func NewActionWeComCustomerTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionWeComCustomerTagLogic {
	return &ActionWeComCustomerTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionWeComCustomerTagLogic) ActionWeComCustomerTag(req *types.ActionCustomerTagRequest) (resp *types.StatusWeComReply, err error) {
	option, err := l.OPT(req)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.PowerX.SCRM.WeCom.ActionWeComCustomerTagRequest(option)

	return &types.StatusWeComReply{
		Status: `success`,
	}, err

}

// OPT
//
//	@Description:
//	@receiver customer
//	@param opt
//	@return option
//	@return err
func (l *ActionWeComCustomerTagLogic) OPT(opt *types.ActionCustomerTagRequest) (option *tagReq.RequestTagMarkTag, err error) {

	if opt.AddTag == nil && opt.RemoveTag == nil {
		return option, errorx.ErrBadRequest
	}
	return &tagReq.RequestTagMarkTag{
		UserID:         opt.UserId,
		ExternalUserID: opt.ExternalUserId,
		AddTag:         opt.AddTag,
		RemoveTag:      opt.RemoveTag,
	}, err
}
