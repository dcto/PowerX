package organization

import (
	"PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/uc/powerx/scrm/wechat/wecom"
	"context"
	"strings"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComDepartMentPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 部门列表/page
func NewListWeComDepartMentPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComDepartMentPageLogic {
	return &ListWeComDepartMentPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComDepartMentPageLogic) ListWeComDepartMentPage(req *types.ListWeComDepartmentReqeust) (resp *types.ListWeComDepartmentReply, err error) {
	data, err := l.svcCtx.PowerX.SCRM.WeCom.FindManyWeComDepartmentsPage(l.ctx, l.OPT(req))

	return &types.ListWeComDepartmentReply{
		List:      l.DTO(data.List),
		PageIndex: data.PageIndex,
		PageSize:  data.PageSize,
		Total:     data.Total,
	}, err

}

// OPT
//
//	@Description:
//	@receiver depart
//	@param opt
//	@return *types.PageOption[wechat.FindManyWeComDepartmentsOption]
func (l *ListWeComDepartMentPageLogic) OPT(opt *types.ListWeComDepartmentReqeust) *types.PageOption[wecom.FindManyWeComDepartmentsOption] {

	option := types.PageOption[wecom.FindManyWeComDepartmentsOption]{
		Option:    wecom.FindManyWeComDepartmentsOption{},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	if v := opt.WeComDepId; v > 0 {
		option.Option.WeComDepId = []int{v}
	}
	if v := opt.Name; v != `` {
		option.Option.Name = v
	}
	option.DefaultPageIfNotSet()

	return &option

}

// DTO
//
//	@Description:
//	@receiver depart
//	@param data
//	@return departments
func (l *ListWeComDepartMentPageLogic) DTO(data []*organization.WeComDepartment) (departments []*types.WeComDepartment) {

	for _, val := range data {
		departments = append(departments, l.dto(val))
	}
	return departments

}

// dto
//
//	@Description:
//	@receiver depart
//	@param val
//	@return *types.WeComDepartment
func (l *ListWeComDepartMentPageLogic) dto(val *organization.WeComDepartment) *types.WeComDepartment {
	var leader []string
	if val.DepartmentLeader != `` {
		leader = strings.Split(val.DepartmentLeader, `,`)
	}
	return &types.WeComDepartment{
		WeComDepId:       val.WeComDepId,
		Name:             val.Name,
		NameEn:           val.NameEn,
		WeComParentId:    val.WeComParentId,
		Order:            val.Order,
		RefDepartmentId:  val.RefDepartmentId,
		DepartmentLeader: leader,
	}
}
