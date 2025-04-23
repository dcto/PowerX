package organization

import (
	"PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/uc/powerx/scrm/wechat/wecom"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 员工列表/page
func NewListWeComUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComUserPageLogic {
	return &ListWeComUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComUserPageLogic) ListWeComUserPage(req *types.ListWeComUserPageReqeust) (resp *types.ListWeComUserPageReply, err error) {
	data, err := l.svcCtx.PowerX.SCRM.WeCom.FindManyWeComUsersPage(l.ctx, l.OPT(req))
	if err != nil {
		logx.Error("list wecom user page failed: %v", err)
		return nil, err
	}
	return &types.ListWeComUserPageReply{
		List:      l.DTO(data.List),
		PageIndex: data.PageIndex,
		PageSize:  data.PageSize,
		Total:     data.Total,
	}, err
}

// OPT
//
//	@Description:
//	@receiver user
//	@param opt
//	@return *types.PageOption[wechat.FindManyWeComUsersOption]
func (l *ListWeComUserPageLogic) OPT(opt *types.ListWeComUserPageReqeust) *types.PageOption[wecom.FindManyWeComUsersOption] {

	option := types.PageOption[wecom.FindManyWeComUsersOption]{
		Option:    wecom.FindManyWeComUsersOption{},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	if opt.Id > 0 {
		option.Option.Ids = []int64{opt.Id}
	}
	if opt.Name != `` {
		option.Option.Names = []string{opt.Name}
	}
	if opt.Alias != `` {
		option.Option.Alias = []string{opt.Alias}
	}
	if opt.Email != `` {
		option.Option.Emails = []string{opt.Email}
	}
	if opt.Mobile != `` {
		option.Option.Mobile = []string{opt.Mobile}
	}
	if opt.OpenUserId != `` {
		option.Option.OpenUserId = []string{opt.OpenUserId}
	}
	if opt.DepartmentId > 0 {
		option.Option.DepartmentId = opt.DepartmentId
	}
	if len(opt.DepartmentIds) > 0 {
		option.Option.DepartmentIds = opt.DepartmentIds
	}
	if opt.Status > 0 {
		option.Option.Status = []int{opt.Status}
	}
	//fmt.Dump(option)
	option.DefaultPageIfNotSet()

	return &option

}

// DTO
//
//	@Description:
//	@receiver user
//	@param data
//	@return users
func (l *ListWeComUserPageLogic) DTO(data []*organization.WeComUser) (users []*types.WeComUser) {

	for _, val := range data {
		users = append(users, l.dto(val))
	}
	return users

}

// dto
//
//	@Description:
//	@receiver user
//	@param val
//	@return *types.WeComUser
func (l *ListWeComUserPageLogic) dto(val *organization.WeComUser) *types.WeComUser {
	return &types.WeComUser{
		UserId:                val.UserId,
		Name:                  val.Name,
		Position:              val.Position,
		Mobile:                val.Mobile,
		Gender:                val.Gender,
		Email:                 val.Email,
		BizMail:               val.BizMail,
		Avatar:                val.Avatar,
		ThumbAvatar:           val.ThumbAvatar,
		Telephone:             val.Telephone,
		Alias:                 val.Alias,
		Address:               val.Address,
		OpenUserId:            val.OpenUserId,
		WeComMainDepartmentId: val.WeComMainDepartmentId,
		Status:                val.Status,
		QrCode:                val.QrCode,
	}
}
