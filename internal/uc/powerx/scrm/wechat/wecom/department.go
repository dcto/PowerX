package wecom

import (
	"PowerX/internal/model/scrm/wechat/wecom/organization"
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
func (uc *WeComUseCase) FindManyWeComDepartmentsPage(ctx context.Context, option *types.PageOption[FindManyWeComDepartmentsOption]) (*types.Page[*organization.WeComDepartment], error) {

	var deps []*organization.WeComDepartment
	var count int64
	query := uc.db.WithContext(ctx).Model(organization.WeComDepartment{})

	if len(option.Option.WeComDepId) > 0 {
		query.Where(`wecom_dep_id in ?`, option.Option.WeComDepId)
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

// GetDepartment
//
// @Description: Retrieve a department and its child nodes based on the specified level.
// @receiver uc
// @param ctx
// @param id Department ID
// @param level The depth of child nodes to retrieve, 0 means only the root node.
// @return *organization.WeComDepartment
// @return error
func (uc *WeComUseCase) GetDepartmentBy(ctx context.Context, id int64, level int) (*organization.WeComDepartment, error) {
	var rootDepartment *organization.WeComDepartment
	// Retrieve the root department
	err := uc.db.WithContext(ctx).
		Where("wecom_dep_id", id).
		First(&rootDepartment).Error
	if err != nil {
		return nil, err
	}

	if level > 0 {
		err = uc.loadChildDepartments(ctx, rootDepartment, level)
		if err != nil {
			return nil, err
		}
	}
	return rootDepartment, nil
}

// loadChildDepartments
//
// @Description: Recursively load child departments up to the specified level.
// @receiver uc
// @param ctx
// @param parent The parent department
// @param level The remaining depth to load child nodes
// @return error
func (uc *WeComUseCase) loadChildDepartments(ctx context.Context, parent *organization.WeComDepartment, level int) error {
	if level == 0 {
		return nil
	}

	var childDepartments []*organization.WeComDepartment
	err := uc.db.WithContext(ctx).
		//Debug().
		Where("wecom_parent_id = ?", parent.WeComDepId).
		Order("wecom_dep_id asc").
		Find(&childDepartments).Error
	if err != nil {
		return err
	}

	parent.Children = childDepartments
	for _, child := range childDepartments {
		err = uc.loadChildDepartments(ctx, child, level-1)
		if err != nil {
			return err
		}
	}

	return nil
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
