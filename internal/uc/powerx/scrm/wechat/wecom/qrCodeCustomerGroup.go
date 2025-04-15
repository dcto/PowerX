package wecom

import (
	"PowerX/internal/model/scene"
	"PowerX/internal/types"
	"strings"
	"time"
)

// CreateWeComCustomerGroupQrcodeRequest
//
//	@Description: 创建客户群活码
//	@receiver this
//	@param opt
//	@return error
func (uc *WeComUseCase) CreateWeComCustomerGroupQrcodeRequest(opt *types.QrcodeActiveRequest) (err error) {

	uc.qrcode.Action(uc.db, []*scene.SceneQRCode{
		{
			QId:                opt.Qid,
			Name:               opt.Name,
			Desc:               opt.Desc,
			Owner:              strings.Join(opt.Owner, `,`),
			RealQrcodeLink:     opt.RealQrcodeLink,
			Platform:           1,
			Classify:           1,
			SceneLink:          opt.SceneLink,
			SafeThresholdValue: opt.SafeThresholdValue,
			ExpiryDate:         opt.ExpiryDate,
			IsAutoActive:       false,
			State:              1,
		},
	})

	return err
}

// UpdateWeComCustomerGroupQrcodeRequest
//
//	@Description: 更新客户群活码
//	@receiver this
//	@param opt
//	@return error
func (uc *WeComUseCase) UpdateWeComCustomerGroupQrcodeRequest(opt *types.QrcodeActiveRequest) (err error) {

	qrcode := uc.qrcode.FindByQid(uc.db, opt.Qid)
	if qrcode != nil {

		qrcode.Name = opt.Name
		qrcode.RealQrcodeLink = opt.RealQrcodeLink
		qrcode.Desc = opt.Desc
		qrcode.Owner = strings.Join(opt.Owner, `,`)
		qrcode.SceneLink = opt.SceneLink
		qrcode.SafeThresholdValue = opt.SafeThresholdValue
		qrcode.ExpiryDate = opt.ExpiryDate
		uc.qrcode.Action(uc.db, []*scene.SceneQRCode{qrcode})

	}

	return err
}

// FindWeComCustomerGroupQrcodePage
//
//	@Description: 客户群活码
//	@receiver this
//	@param opt
//	@return reply
//	@return error
func (uc *WeComUseCase) FindWeComCustomerGroupQrcodePage(option *types.PageOption[types.ListWeComGroupQrcodeActiveReqeust]) (reply *types.Page[*scene.SceneQRCode], err error) {

	var code []*scene.SceneQRCode
	var count int64
	query := uc.db.WithContext(uc.ctx).Model(scene.SceneQRCode{}).Where(`state < 3`)

	if v := option.Option.Name; v != `` {
		query.Where("name like ?", "%"+v+"%")
	}
	if v := option.Option.Qid; v != `` {
		query.Where("qid = ?", v)
	}
	if v := option.Option.UserId; v != `` {
		query.Where("POSITION(? IN owner) > 0", v)
	}
	if v := option.Option.State; v > 0 {
		query.Where("state = ?", v)
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if option.PageIndex != 0 && option.PageSize != 0 {
		query.Offset((option.PageIndex - 1) * option.PageSize).Order(`expiry_date ASC`).Limit(option.PageSize)
	}
	_ = query.Find(&code).Error

	return &types.Page[*scene.SceneQRCode]{
		List:      code,
		PageIndex: option.PageIndex,
		PageSize:  option.PageSize,
		Total:     count,
	}, err
}

// ActionCustomerGroupQrcode
//
//	@Description:
//	@receiver this
//	@param qid
//	@return error
func (uc *WeComUseCase) ActionCustomerGroupQrcode(qid string, action int) error {
	column := make(map[string]interface{})
	column[`state`] = action
	if action == 3 {
		column[`deleted_at`] = time.Now()
	}
	uc.modelWeComQrcode.qrcode.UpdateColumn(uc.db, qid, column)

	return nil
}

// UpdateSceneQRCodeLink
//
//	@Description:
//	@receiver this
//	@param qid
//	@param link
//	@return error
func (uc *WeComUseCase) UpdateSceneQRCodeLink(qid string, link string) error {

	column := make(map[string]interface{})
	column[`active_qrcode_link`] = link
	uc.modelWeComQrcode.qrcode.UpdateColumn(uc.db, qid, column)

	return nil
}
