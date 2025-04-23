package customer

import (
	"context"
	"github.com/ArtisanCloud/PowerSocialite/v3/src/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncWeComCustomerOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量同步客户信息(根据员工ID同步/节流)
func NewSyncWeComCustomerOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWeComCustomerOptionLogic {
	return &SyncWeComCustomerOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWeComCustomerOptionLogic) SyncWeComCustomerOption(req *types.WeComCustomersRequest) (resp *types.WeComListCustomersReply, err error) {
	data, err := l.svcCtx.PowerX.SCRM.WeCom.PullListWeComCustomerRequest(req.UserId)

	return &types.WeComListCustomersReply{
		List: l.DTO(data),
	}, err
}

// DTO
//
//	@Description:
//	@receiver cMsg
//	@param data
//	@return resp
func (l *SyncWeComCustomerOptionLogic) DTO(data []*response.ResponseExternalContact) (resp []*types.WeComCustomer) {

	for _, obj := range data {
		resp = append(resp, l.dto(obj))
	}
	return resp

}

// dto
//
//	@Description:
//	@receiver cMsg
//	@param contact
//	@return *types.WeComCustomer
func (l *SyncWeComCustomerOptionLogic) dto(contact *response.ResponseExternalContact) *types.WeComCustomer {
	return &types.WeComCustomer{
		ExternalContact: l.contact(contact.ExternalContact),
		FollowUser:      l.follow(contact.FollowInfo),
		NextCursor:      ``,
	}
}

// contact
//
//	@Description:
//	@receiver cMsg
//	@param data
//	@return types.WeComCustomersWithExternalContactExternalProfile
func (l *SyncWeComCustomerOptionLogic) contact(data *models.ExternalContact) types.WeComCustomersWithExternalContactExternalProfile {
	return types.WeComCustomersWithExternalContactExternalProfile{
		ExternalUserId: data.ExternalUserID,
		Name:           data.Name,
		Position:       data.Position,
		Avatar:         data.Avatar,
		CorpName:       data.CorpName,
		CorpFullName:   data.CorpFullName,
		Type:           data.Type,
		Gender:         data.Gender,
		UnionId:        data.UnionID,
		ExternalProfile: types.ExternalContactExternalProfileWithExternalProfile{
			l.contactExternalProfile(data.ExternalProfile),
		},
	}
}

// follow
//
//	@Description:
//	@receiver cMsg
//	@param follow
//	@return *types.WeComCustomersWithFollowUser
func (l *SyncWeComCustomerOptionLogic) follow(follow *models.FollowUser) *types.WeComCustomersWithFollowUser {

	if follow == nil {
		return nil
	}
	return &types.WeComCustomersWithFollowUser{
		UserId:         follow.UserID,
		Remark:         follow.Remark,
		Description:    follow.Description,
		CreatedTime:    follow.CreateTime,
		Tags:           nil,
		WeComChannels:  l.followWeComChannels(follow.WechatChannels),
		RemarkCorpName: follow.RemarkCorpName,
		RemarkMobiles:  follow.RemarkMobiles,
		OpenUserId:     follow.OperUserID,
		AddWay:         follow.AddWay,
		State:          follow.State,
	}

}

// contactExternalProfile
//
//	@Description:
//	@receiver cMsg
//	@param profiles
//	@return externalProfile
func (l *SyncWeComCustomerOptionLogic) contactExternalProfile(profiles *models.ExternalProfile) (externalProfile []*types.ExternalContactExternalProfileExternalProfileWithExternalAttr) {

	if profiles != nil {
		for _, obj := range profiles.ExternalAttr {
			externalProfile = append(externalProfile, &types.ExternalContactExternalProfileExternalProfileWithExternalAttr{
				Type: obj.Type,
				Name: obj.Name,
				Text: types.ExternalContactExternalProfileExternalProfileExternalAttrWithText{obj.Text.Value},
				Web:  types.ExternalContactExternalProfileExternalProfileExternalAttrWithWeb{obj.Web.URL, obj.Web.Title},
				Miniprogram: types.ExternalContactExternalProfileExternalProfileExternalAttrWithMiniprogram{
					obj.MiniProgram.AppID,
					obj.MiniProgram.PagePath,
					obj.MiniProgram.Title,
				},
			})
		}
	}

	return externalProfile
}

// followWeComChannels
//
//	@Description:
//	@receiver cMsg
//	@param channel
//	@return data
func (l *SyncWeComCustomerOptionLogic) followWeComChannels(channel *models.WechatChannel) (data types.WeComCustomersFollowUserWithWeComChannels) {
	if channel != nil {
		data.Nickname = channel.NickName
		data.Source = channel.Source
	}
	return data
}
