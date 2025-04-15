package scene

import (
	"context"
	"fmt"
	"strings"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailQRCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailQRCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailQRCodeLogic {
	return &DetailQRCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DetailQRCode
//
//	@Description:
//	@receiver qrcode
//	@param opt
//	@return resp
//	@return err
func (qrcode *DetailQRCodeLogic) DetailQRCode(opt *types.SceneRequest) (resp *types.SceneQRCodeActiveReply, err error) {
	if opt.Qid == `` {
		return nil, fmt.Errorf(`Qid error`)
	}

	detail := qrcode.svcCtx.PowerX.Scene.Scene.FindOneSceneQRCodeDetail(opt.Qid)
	go qrcode.svcCtx.PowerX.Scene.Scene.IncreaseSceneCpaNumber(opt.Qid)

	return &types.SceneQRCodeActiveReply{
		QId:                detail.QId,
		Name:               detail.Name,
		Desc:               detail.Desc,
		Owner:              strings.Split(detail.Owner, `,`),
		RealQRCodeLink:     detail.RealQRCodeLink,
		Platform:           detail.Platform,
		Classify:           detail.Classify,
		SceneLink:          detail.SceneLink,
		SafeThresholdValue: detail.SafeThresholdValue,
		ExpiryDate:         detail.ExpiryDate,
		State:              detail.State,
		ActiveQRCodeLink:   detail.ActiveQRCodeLink,
		CPA:                detail.Cpa,
	}, err
}
