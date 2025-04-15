package qrcode

import (
	"PowerX/internal/model/scene"
	"context"
	"strings"
	"time"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComQRCodePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 场景码列表/page
func NewListWeComQRCodePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComQRCodePageLogic {
	return &ListWeComQRCodePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComQRCodePageLogic) ListWeComQRCodePage(req *types.ListWeComGroupQRCodeActiveReqeust) (resp *types.ListWeComQRCodeActiveReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.FindWeComCustomerGroupQRCodePage(l.OPT(req))

	return &types.ListWeComQRCodeActiveReply{
		List:      l.DTO(reply),
		PageIndex: reply.PageIndex,
		PageSize:  reply.PageSize,
		Total:     reply.Total,
	}, err
}

// @Description:
// @receiver qrcode
// @param opt
// @return *types.PageOption[types.ListWeComGroupQRCodeActiveReqeust]
func (l *ListWeComQRCodePageLogic) OPT(opt *types.ListWeComGroupQRCodeActiveReqeust) *types.PageOption[types.ListWeComGroupQRCodeActiveReqeust] {

	option := types.PageOption[types.ListWeComGroupQRCodeActiveReqeust]{
		Option:    types.ListWeComGroupQRCodeActiveReqeust{},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	if v := opt.UserId; v != `` {
		option.Option.UserId = v
	}
	if v := opt.Name; v != `` {
		option.Option.Name = v
	}
	if v := opt.Qid; v != `` {
		option.Option.Qid = v
	}
	if v := opt.State; v > 0 {
		option.Option.State = v
	}
	return &option

}

// DTO
//
//	@Description:
//	@receiver qrcode
//	@param data
//	@return reply
func (l *ListWeComQRCodePageLogic) DTO(data *types.Page[*scene.SceneQRCode]) (reply []*types.WeComQRCodeActive) {

	if data.List != nil {
		for _, obj := range data.List {
			reply = append(reply, l.dto(obj))
		}
	}

	return reply
}

// dto
//
//	@Description:
//	@receiver qrcode
//	@param obj
//	@return *types.WeComQRCodeActive
func (l *ListWeComQRCodePageLogic) dto(obj *scene.SceneQRCode) *types.WeComQRCodeActive {

	return &types.WeComQRCodeActive{
		QId:                obj.QId,
		Name:               obj.Name,
		Desc:               obj.Desc,
		Owner:              strings.Split(obj.Owner, `,`),
		RealQRCodeLink:     obj.RealQRCodeLink,
		Platform:           obj.Platform,
		Classify:           obj.Classify,
		SceneLink:          obj.SceneLink,
		SafeThresholdValue: obj.SafeThresholdValue,
		ExpiryDate:         obj.ExpiryDate,
		ExpiryState:        int(obj.ExpiryDate - time.Now().Unix()),
		ActiveQRCodeLink:   obj.ActiveQRCodeLink,
		CPA:                obj.Cpa,
		State:              obj.State,
	}
}
