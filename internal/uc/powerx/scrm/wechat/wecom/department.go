package wecom

import (
	"PowerX/internal/model/scrm/organization"
	"PowerX/internal/types"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/request"
)

// CreateWeComDepartmentRequest
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@param dep
//	@return error
func (uc *WeComUseCase) CreateWeComDepartmentRequest(ctx context.Context, dep *organization.WeComDepartment) (err error) {

	create, err := uc.Client.Department.Create(ctx, &request.RequestDepartmentInsert{
		Name:     dep.Name,
		NameEn:   dep.NameEn,
		ParentID: dep.WeComParentId,
		Order:    dep.Order,
		ID:       int(dep.Id),
	})

	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.create.wecom.department.error`, create.ResponseWork)

	}
	return err

}

// UpdateWeComDepartmentRequest
//
//	@Description:
//	@receiver this
//	@param ctx
//	@param dep
//	@return err
func (uc *WeComUseCase) UpdateWeComDepartmentRequest(ctx context.Context, dep *organization.WeComDepartment) (err error) {

	update, err := uc.Client.Department.Update(ctx, &request.RequestDepartmentUpdate{
		Name:     dep.Name,
		NameEn:   dep.NameEn,
		ParentID: dep.WeComParentId,
		Order:    dep.Order,
		ID:       int(dep.Id),
	})
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.update.wecom.department.error`, update.ResponseWork)

	}

	return err

}

// FindManyWeComDepartmentsPage
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@param option
//	@return *types.Page[*organization.WeComDepartment]
//	@return error
func (uc *WeComUseCase) FindManyWeComDepartmentsPage(ctx context.Context, option *types.PageOption[FindManyWechatDepartmentsOption]) (*types.Page[*organization.WeComDepartment], error) {

	var deps []*organization.WeComDepartment
	var count int64
	query := uc.db.WithContext(ctx).Model(organization.WeComDepartment{})

	if len(option.Option.WeComDepId) > 0 {
		query.Where(`we_work_dep_id in ?`, option.Option.WeComDepId)
	}

	if v := option.Option.Name; v == `` {
		query.Where("name like ?", "%"+v+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if option.PageIndex != 0 && option.PageSize != 0 {
		query.Offset((option.PageIndex - 1) * option.PageSize).Limit(option.PageSize)
	}
	err := query.Find(&deps).Error

	return &types.Page[*organization.WeComDepartment]{
		List:      deps,
		PageIndex: option.PageIndex,
		PageSize:  option.PageSize,
		Total:     count,
	}, err

}

// FindAllWechatDepartments
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@return departments
//	@return err
func (uc *WeComUseCase) FindAllWechatDepartments(ctx context.Context) (departments []*organization.WeComDepartment, err error) {

	err = uc.db.WithContext(ctx).Find(&departments).Error
	return departments, err

}
