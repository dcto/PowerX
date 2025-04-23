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

type ListWeComDepartmentPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 部门列表/page
func NewListWeComDepartmentPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComDepartmentPageLogic {
	return &ListWeComDepartmentPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComDepartmentPageLogic) ListWeComDepartmentPage(req *types.ListWeComDepartmentPageReqeust) (resp *types.ListWeComDepartmentPageReply, err error) {
	data, err := l.svcCtx.PowerX.SCRM.WeCom.FindManyWeComDepartmentsPage(l.ctx, l.OPT(req))

	return &types.ListWeComDepartmentPageReply{
		List:      TransformWeComDepartmentsToReply(data.List),
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
func (l *ListWeComDepartmentPageLogic) OPT(opt *types.ListWeComDepartmentPageReqeust) *types.PageOption[wecom.FindManyWeComDepartmentsOption] {

	option := types.PageOption[wecom.FindManyWeComDepartmentsOption]{
		Option:    wecom.FindManyWeComDepartmentsOption{},
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}
	if v := opt.DepartmentId; v > 0 {
		option.Option.WeComDepId = []int{v}
	}
	if v := opt.Name; v != `` {
		option.Option.Name = v
	}
	option.DefaultPageIfNotSet()

	return &option

}

func TransformWeComDepartmentsToReply(data []*organization.WeComDepartment) (departments []*types.WeComDepartment) {
	if data == nil || len(data) == 0 {
		return nil
	}

	for _, val := range data {
		departments = append(departments, TransformWeComDepartmentToReply(val))
	}
	return departments

}

func TransformWeComDepartmentToReply(val *organization.WeComDepartment) *types.WeComDepartment {
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
		Children:         TransformWeComDepartmentsToReply(val.Children),
	}
}
