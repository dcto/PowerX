package customer

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComCustomerGroupLimitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户群列表/limit
func NewListWeComCustomerGroupLimitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComCustomerGroupLimitLogic {
	return &ListWeComCustomerGroupLimitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComCustomerGroupLimitLogic) ListWeComCustomerGroupLimit(req *types.WeComCustomerGroupRequest) (resp *types.WeComListCustomerGroupReply, err error) {
	newMap, _ := power.StructToHashMap(req.OwnerFilter)
	option := &request.RequestGroupChatList{
		StatusFilter: req.StatusFilter,
		OwnerFilter:  newMap,
		Cursor:       req.Cursor,
		Limit:        1000,
	}
	if option.Limit == 0 {
		option.Limit = req.Limit
	}
	list, err := l.svcCtx.PowerX.SCRM.WeCom.PullListWeComCustomerGroupRequest(option)

	if list != nil {
		resp = l.DTO(list)
	}
	return resp, err
}

// DTO
//
//	@Description:
//	@receiver cGroup
//	@param data
//	@return *types.WeComListCustomerGroupReply
func (l *ListWeComCustomerGroupLimitLogic) DTO(data []*response.ResponseGroupChatGet) *types.WeComListCustomerGroupReply {

	reply := types.WeComListCustomerGroupReply{}
	for _, obj := range data {
		if obj != nil {
			reply.List = append(reply.List, l.dto(obj.GroupChat))
		}
	}

	return &reply

}

// dto
//
//	@Description:
//	@receiver cGroup
//	@param chat
//	@return types.WeComCustomerGroup
func (l *ListWeComCustomerGroupLimitLogic) dto(chat *response.GroupChat) types.WeComCustomerGroup {

	return types.WeComCustomerGroup{
		ChatId:     chat.ChatID,
		Name:       chat.Name,
		Owner:      chat.Owner,
		CreateTime: chat.CreateTime,
		Notice:     chat.Notice,
		MemberList: l.members(chat.MemberList),
		AdminList:  l.admins(chat.AdminList),
	}
}

// members
//
//	@Description:
//	@receiver cGroup
//	@param members
//	@return list
func (l *ListWeComCustomerGroupLimitLogic) members(members []*response.Member) (list []*types.WeComCustomerGroupMemberList) {

	for _, val := range members {
		list = append(list, &types.WeComCustomerGroupMemberList{
			UserId:        val.UserID,
			Type:          val.Type,
			JoinTime:      val.JoinTime,
			JoinScene:     val.JoinScene,
			Invitor:       l.weComCustomerGroupMemberListInvitor(val.Invitor),
			GroupNickname: val.GroupNickname,
			Name:          val.Name,
			UnionId:       val.UnionID,
		})
	}
	return list

}

// admins
//
//	@Description:
//	@receiver cGroup
//	@param admins
//	@return list
func (l *ListWeComCustomerGroupLimitLogic) admins(admins []*response.Admin) (list []*types.WeComCustomerGroupAdminList) {

	for _, val := range admins {
		list = append(list, &types.WeComCustomerGroupAdminList{
			UserId: val.UserID,
		})
	}
	return list

}

// weComCustomerGroupMemberListInvitor
//
//	@Description:
//	@receiver cGroup
//	@param invitor
//	@return info
func (l *ListWeComCustomerGroupLimitLogic) weComCustomerGroupMemberListInvitor(invitor *response.Invitor) (info types.WeComCustomerGroupMemberListInvitor) {
	if invitor != nil {
		info.UserId = invitor.UserID
	}
	return info
}
