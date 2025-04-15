package customer

import (
	customer2 "PowerX/internal/model/scrm/wechat/wecom/customer"
	"PowerX/internal/uc/powerx/scrm/wechat/wecom"
	"context"
	"encoding/json"
	"strings"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComCustomerPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 所有客户列表/page
func NewListWeComCustomerPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComCustomerPageLogic {
	return &ListWeComCustomerPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComCustomerPageLogic) ListWeComCustomerPage(req *types.WeComCustomersRequest) (resp *types.WeComListCustomersReply, err error) {
	data, err := l.svcCtx.PowerX.SCRM.WeCom.FindManyWeComCustomerPage(l.ctx, l.OPT(req), req.Sync)
	return &types.WeComListCustomersReply{
		List:      l.DTO(data.List),
		PageIndex: data.PageIndex,
		PageSize:  data.PageSize,
		Total:     data.Total,
	}, err

}

// OPT
//
//	@Description:
//	@receiver customer
//	@param opt
//	@return *types.PageOption[wechat.FindManyWeComCustomerOption]
func (l *ListWeComCustomerPageLogic) OPT(opt *types.WeComCustomersRequest) *types.PageOption[wecom.FindManyWeComCustomerOption] {

	option := types.PageOption[wecom.FindManyWeComCustomerOption]{
		Option: wecom.FindManyWeComCustomerOption{
			UserId: opt.UserId,
			Name:   opt.Name,
			TagId:  opt.TagId,
		},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}

	option.DefaultPageIfNotSet()

	return &option

}

// DTO
//
//	@Description:
//	@receiver customer
//	@param data
//	@return resp
func (l *ListWeComCustomerPageLogic) DTO(data []*customer2.WeComExternalContact) (resp []*types.WeComCustomer) {

	if data != nil {
		for _, obj := range data {
			resp = append(resp, l.dto(obj))
		}
	}
	return resp

}

// dto
//
//	@Description:
//	@receiver customer
//	@param contact
//	@return *types.WeComCustomer
func (l *ListWeComCustomerPageLogic) dto(contact *customer2.WeComExternalContact) *types.WeComCustomer {

	return &types.WeComCustomer{
		ExternalContact: types.WeComCustomersWithExternalContactExternalProfile{

			ExternalUserId:  contact.ExternalUserId,
			Name:            contact.Name,
			Position:        contact.Position,
			Avatar:          contact.Avatar,
			CorpName:        contact.CorpName,
			CorpFullName:    contact.CorpFullName,
			Type:            int(contact.WXType),
			Gender:          contact.Gender,
			UnionId:         contact.UnionId,
			UserId:          contact.UserId,
			ExternalProfile: l.externalContactExternalProfileWithExternalProfile(contact.ExternalProfile),
		},
		FollowUser: l.weComCustomersWithFollowUser(&contact.WeComExternalContactFollow),
		NextCursor: ``,
	}
}

// externalContactExternalProfileWithExternalProfile
//
//	@Description:
//	@receiver customer
//	@param attr
//	@return data
func (l *ListWeComCustomerPageLogic) externalContactExternalProfileWithExternalProfile(attr string) (data types.ExternalContactExternalProfileWithExternalProfile) {
	if attr != `` {
		_ = json.Unmarshal([]byte(attr), &data)
	}
	return data
}

// weComCustomersWithFollowUser
//
//	@Description:
//	@receiver customer
//	@param follow
//	@return *types.WeComCustomersWithFollowUser
func (l *ListWeComCustomerPageLogic) weComCustomersWithFollowUser(follow *customer2.WeComExternalContactFollow) *types.WeComCustomersWithFollowUser {

	if follow == nil {
		return nil
	}
	var tags []types.WeComCustomersFollowUserWithTags
	if follow.Tags != `` {
		_ = json.Unmarshal([]byte(follow.Tags), &tags)
	}

	return &types.WeComCustomersWithFollowUser{
		UserId:         follow.UserId,
		Remark:         follow.Remark,
		Description:    follow.Description,
		CreatedTime:    follow.CreatedTime,
		Tags:           tags,
		TagIds:         strings.Split(follow.TagIds, `,`),
		WeComChannels:  l.WeComCustomersFollowUserWithWeComChannels(follow.WeComChannels),
		RemarkCorpName: follow.RemarkCorpName,
		RemarkMobiles:  []string{follow.RemarkMobiles},
		OpenUserId:     follow.OpenUserId,
		AddWay:         follow.AddWay,
		State:          follow.State,
	}
}

// WeComCustomersFollowUserWithWeComChannels
//
//	@Description:
//	@receiver customer
//	@param channels
//	@return channel
func (l *ListWeComCustomerPageLogic) WeComCustomersFollowUserWithWeComChannels(channels string) (channel types.WeComCustomersFollowUserWithWeComChannels) {

	if channels == `` {
		return channel
	}
	_ = json.Unmarshal([]byte(channels), &channel)

	return channel

}
