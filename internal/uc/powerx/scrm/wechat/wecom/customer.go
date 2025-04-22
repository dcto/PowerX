package wecom

import (
	"PowerX/internal/model/powerModel"
	customer2 "PowerX/internal/model/scrm/wechat/wecom/customer"
	"PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerSocialite/v3/src/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

// FindManyWeComCustomerPage
//
//	@Description:
//	@receiver this
//	@param ctx
//	@param opt
//	@param sync
//	@return *types.Page[*customer.WeComExternalContact]
//	@return error
func (uc *WeComUseCase) FindManyWeComCustomerPage(ctx context.Context, opt *types.PageOption[FindManyWeComCustomerOption], sync int) (*types.Page[*customer2.WeComExternalContact], error) {

	if sync > 0 {
		uc.pullSyncWeComCustomerRequest(opt.Option.UserId)
	}

	var customers []*customer2.WeComExternalContact
	var count int64
	query := uc.db.WithContext(ctx).Table(new(customer2.WeComExternalContact).TableName() + ` AS a`).
		Joins(fmt.Sprintf(`LEFT JOIN %s AS b ON a.external_user_id=b.external_user_id`, new(customer2.WeComExternalContactFollow).TableName()))

	if opt.PageIndex == 0 {
		opt.PageIndex = 1
	}
	if opt.PageSize == 0 {
		opt.PageSize = powerModel.PageDefaultSize
	}
	query = buildFindManyCustomerQueryNoPage(query, &opt.Option)

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Preload(`WeComExternalContactFollow`).Find(&customers).Error

	return &types.Page[*customer2.WeComExternalContact]{
		List:      customers,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, err
}

// buildFindManyCustomerQueryNoPage
//
//	@Description:
//	@param query
//	@param opt
//	@return *gorm.DB
func buildFindManyCustomerQueryNoPage(query *gorm.DB, opt *FindManyWeComCustomerOption) *gorm.DB {

	if v := opt.UserId; v != `` {
		query.Where("a.user_id = ?", v)
	}

	if v := opt.Name; v != `` {
		query.Where("a.name like ?", "%"+v+"%")
	}

	if v := opt.TagId; v != `` {
		query.Where(`POSITION(? IN b.tag_ids) > 0`, v)
	}
	return query
}

// PullListWeComCustomerRequest
//
//	@Description:
//	@receiver this
//	@param userID
//	@return []*response.ResponseExternalContact
//	@return error
func (uc *WeComUseCase) PullListWeComCustomerRequest(userID ...string) ([]*response.ResponseExternalContact, error) {

	var err error

	// 外部联系人和客户                                     ExternalContact * externalContact.Client
	// 客户群                                             ExternalContactGroupChat * groupChat.Client
	// 外部联系人和客户                                     ExternalContactContactWay * contactWay.Client
	// 规则                                               ExternalContactCustomerStrategy * customerStrategy.Client
	// 联系客户统计                                        ExternalContactStatistics * statistics.Client
	// 欢迎语                                             ExternalContactGroupWelcomeTemplate * groupWelcomeTemplate.Client
	// 学校                                               ExternalContactSchool * school.Client
	// 朋友圈/发表                                         ExternalContactMoment * moment.Client
	// 规则组                                             ExternalContactMomentStrategy * momentStrategy.Client
	// 企业群发                                            ExternalContactMessageTemplate * messageTemplate.Client
	// 企业标签库                                          ExternalContactTag * tag2.Client
	// 人事变动                                            ExternalContactTransfer * transfer.Client

	info, err := uc.Client.ExternalContact.BatchGet(uc.ctx, userID, ``, 1000)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.customer.list.error`, info.ResponseWork)

	}
	contacts := []customer2.WeComExternalContact{}
	follows := []customer2.WeComExternalContactFollow{}

	for _, val := range info.ExternalContactList {
		contacts = append(contacts, transferExternalContactToModel(val.ExternalContact, val.FollowInfo.UserID))
		follows = append(follows, transferExternalContactFollowToModel(val.FollowInfo, val.ExternalContact.ExternalUserID))
	}
	err = uc.db.Clauses(
		clause.OnConflict{Columns: []clause.Column{{Name: `external_user_id`}}, UpdateAll: true}).CreateInBatches(&contacts, 100).Error
	err = uc.db.Clauses(
		clause.OnConflict{Columns: []clause.Column{{Name: `external_user_id`}}, UpdateAll: true}).CreateInBatches(&follows, 100).Error
	if err != nil {
		logx.Errorf(`scrm.wecom.customer.contract.error. %v`, err)
	}
	if info != nil {
		return info.ExternalContactList, nil
	}

	return nil, err

}

// transferExternalContactToModel
//
//	@Description:
//	@param contact
//	@return *customer.WeComExternalContact
func transferExternalContactToModel(contact *models.ExternalContact, userID string) customer2.WeComExternalContact {
	return customer2.WeComExternalContact{

		ExternalUserId:  contact.ExternalUserID,
		AppId:           ``,
		CorpId:          ``,
		OpenId:          ``,
		UnionId:         contact.UnionID,
		UserId:          userID,
		Name:            contact.Name,
		Mobile:          ``,
		Position:        contact.Position,
		Avatar:          contact.Avatar,
		CorpName:        contact.CorpName,
		CorpFullName:    contact.CorpFullName,
		ExternalProfile: ``,
		Gender:          contact.Gender,
		WXType:          int8(contact.Type),
		Status:          1,
		Active:          true,
	}
}

// transferExternalContactFollowToModel
//
//	@Description:
//	@param follow
//	@param externalUserID
//	@return customer.WeComExternalContactFollow
func transferExternalContactFollowToModel(follow *models.FollowUser, externalUserID string) customer2.WeComExternalContactFollow {

	tags, _ := json.Marshal(follow.Tags)
	remarkMobiles, _ := json.Marshal(follow.RemarkMobiles)
	return customer2.WeComExternalContactFollow{
		ExternalUserId: externalUserID,
		UserId:         follow.UserID,
		Remark:         follow.Remark,
		Description:    follow.Description,
		CreatedTime:    follow.CreateTime,
		Tags:           string(tags),
		TagIds:         strings.Join(follow.TagIDs, `,`),
		WeComChannels:  string(remarkMobiles),
		RemarkCorpName: follow.RemarkCorpName,
		RemarkMobiles:  ``,
		OpenUserId:     follow.OperUserID,
		AddWay:         follow.AddWay,
		State:          follow.State,
	}
}

// pullSyncWeComCustomer
//
//	@Description: 全量/增量同步客户信息
//	@receiver this
//	@param ids
func (uc *WeComUseCase) pullSyncWeComCustomerRequest(ids ...string) {

	if len(ids) > 0 && ids[0] == `` {
		workUsers := uc.modelWeComOrganization.user.Query(uc.db)
		ids = organization.AdapterUserSliceUserIDs(func(users []*organization.WeComUser) (ids []string) {
			for _, user := range users {
				ids = append(ids, user.UserId)
			}
			return ids
		})(workUsers)
	}

	_, _ = uc.PullListWeComCustomerRequest(ids...)

}
