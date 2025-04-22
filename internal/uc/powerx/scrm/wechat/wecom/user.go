package wecom

import (
	organization2 "PowerX/internal/model/organization"
	"PowerX/internal/model/powerModel"
	organization3 "PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/types"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/request"
	"gorm.io/gorm"
	"strings"
)

// CreateWeComUser
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@param dep
//	@return error
func (uc *WeComUseCase) CreateWeComUser(ctx context.Context, user *organization3.WeComUser) (err error) {

	create, err := uc.Client.User.Create(ctx, uc.userModelToWeComRequest(user))
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.create.wecom.user.error`, *create)
	}

	if err == nil {
		uc.modelWeComOrganization.user.BatchUpsert(uc.db, []*organization3.WeComUser{user})
	}

	return err

}

// UpdateWeComUser
//
//	@Description:
//	@receiver this
//	@param ctx
//	@param dep
//	@return err
func (uc *WeComUseCase) UpdateWeComUser(ctx context.Context, user *organization3.WeComUser) (err error) {

	update, err := uc.Client.User.Update(ctx, uc.userModelToWeComRequest(user))

	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.update.wecom.organization.user.error`, *update)
	}

	if err == nil {
		uc.modelWeComOrganization.user.BatchUpsert(uc.db, []*organization3.WeComUser{user})
	}
	return err

}

// userModelToWeComRequest
//
//	@Description:
//	@param user
//	@return *request.RequestUserDetail
func (uc *WeComUseCase) userModelToWeComRequest(user *organization3.WeComUser) *request.RequestUserDetail {

	return &request.RequestUserDetail{
		Userid:         user.UserId,
		Name:           user.Name,
		Alias:          user.Alias,
		Mobile:         user.Mobile,
		Position:       user.Position,
		Email:          user.Email,
		BizMail:        user.BizMail,
		Telephone:      user.Telephone,
		Address:        user.Address,
		MainDepartment: user.WeComMainDepartmentId,
	}
}

// PullSyncDepartmentsAndUsers
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@return error
func (uc *WeComUseCase) PullSyncDepartmentsAndUsers(ctx context.Context) error {

	list, err := uc.Client.Department.SimpleList(ctx, 1)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.sync.organization.list.error`, list.ResponseWork)
	}

	if err != nil {
		return err
	}

	uc.gLock.Add(len(list.DepartmentIDs))
	for _, val := range list.DepartmentIDs {
		go func(val response.DepartmentID) {
			defer uc.gLock.Done()
			uc.syncDepartment(ctx, val)
			uc.syncDepartmentUsers(ctx, val)

		}(val)

	}
	uc.gLock.Wait()
	return err
}

// syncDepartment
//
//	@Description:
//	@receiver this
//	@param val
func (uc *WeComUseCase) syncDepartment(ctx context.Context, val response.DepartmentID) {

	department, err := uc.Client.Department.Get(ctx, val.ID)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wechat.sync.organization.department.error`, department.ResponseWork)
	}

	if err == nil && department.Department != nil {

		uc.modelWeComOrganization.department.Action(uc.db, []*organization3.WeComDepartment{
			{
				WeComDepId:       department.Department.ID,
				Name:             department.Department.Name,
				NameEn:           department.Department.NameEN,
				WeComParentId:    department.Department.ParentID,
				Order:            department.Department.Order,
				DepartmentLeader: strings.Join(department.Department.DepartmentLeaders, `,`),
			},
		})

	}

}

// user
//
//	@Description:
//	@receiver this
//	@param val
func (uc *WeComUseCase) syncDepartmentUsers(ctx context.Context, val response.DepartmentID) {

	resUsers, err := uc.Client.User.GetDetailedDepartmentUsers(ctx, val.ID, 0)
	//fmt.Dump(val, resUsers)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wecom.sync.organization.user.error`, resUsers.ResponseWork)
	}

	if err == nil && len(resUsers.UserList) > 0 {
		users := []*organization3.WeComUser{}
		for _, employee := range resUsers.UserList {
			if employee != nil {
				open, err := uc.Client.User.UserIdToOpenID(ctx, employee.UserID)
				if err != nil {
					panic(err)
				} else {
					err = uc.help.error(`scrm.wecom.sync.organization.user.error`, open.ResponseWork)
				}
				users = append(users, &organization3.WeComUser{
					UserId:                employee.UserID,
					Name:                  employee.Name,
					Position:              employee.Position,
					Mobile:                employee.Mobile,
					Email:                 employee.Email,
					Alias:                 employee.Alias,
					OpenUserId:            open.OpenID,
					WeComMainDepartmentId: employee.MainDepartment,
					Status:                employee.Status,
					QrCode:                employee.QrCode,
					RefUserId:             0,
				})
			}
		}
		//fmt.Dump(users)
		uc.modelWeComOrganization.user.BatchUpsert(uc.db, users)
		// sync to local
		//uc.modelOrganization.user.Action(uc.db, uc.userFromWeComSyncToLocal(users))

	}

}

// buildFindManyUsersQueryNoPage
//
//	@Description:
//	@param query
//	@param opt
//	@return *gorm.DB
func buildFindManyUsersQueryNoPage(query *gorm.DB, opt *FindManyWeComUsersOption) *gorm.DB {
	if len(opt.Ids) > 0 {
		query.Where("id in ?", opt.Ids)
	}
	if len(opt.Names) > 0 {
		query.Where("name in ?", opt.Names)
	}
	if len(opt.Emails) > 0 {
		query.Where("email in ?", opt.Emails)
	}
	if len(opt.Mobile) > 0 {
		query.Where("mobile in ?", opt.Mobile)
	}
	if len(opt.Alias) > 0 {
		query.Where("alias in ?", opt.Alias)
	}
	if len(opt.OpenUserId) > 0 {
		query.Where("open_user_id in ?", opt.OpenUserId)
	}
	if len(opt.WeComMainDepartmentId) > 0 {
		query.Where("wecom_main_department_id in ? ", opt.WeComMainDepartmentId)
	}
	if len(opt.Status) > 0 {
		query.Where("status in ?", opt.Status)
	}
	return query
}

// FindManyWeComUsersPage
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@param opt
//	@return *types.Page[*organization.WeComUser]
//	@return error
func (uc *WeComUseCase) FindManyWeComUsersPage(ctx context.Context, opt *types.PageOption[FindManyWeComUsersOption]) (*types.Page[*organization3.WeComUser], error) {

	var users []*organization3.WeComUser
	var count int64

	query := uc.db.WithContext(ctx).Table(uc.modelWeComOrganization.user.TableName())

	if opt.PageIndex == 0 {
		opt.PageIndex = 1
	}
	if opt.PageSize == 0 {
		opt.PageSize = powerModel.PageDefaultSize
	}
	query = buildFindManyUsersQueryNoPage(query, &opt.Option)

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Find(&users).Error

	return &types.Page[*organization3.WeComUser]{
		List:      users,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, err
}

// getWeComUserIDs
//
//	@Description:
//	@receiver uc
//	@param ctx
//	@param opt
//	@return *types.Page[*organization.WeComUser]
//	@return error
func (uc *WeComUseCase) getWeComUserIDs(ctx context.Context) (ids []string, err error) {

	ids = organization3.AdapterUserSliceUserIDs(func(users []*organization3.WeComUser) (ids []string) {
		for _, user := range users {
			ids = append(ids, user.UserId)
		}
		return ids
	})(uc.modelWeComOrganization.user.Query(uc.db))

	return ids, err

}

// userFromWeComSyncToLocal
//
//	@Description:
//	@receiver this
//	@param fromUser
//	@return toUser
func (uc *WeComUseCase) userFromWeComSyncToLocal(fromUser []*organization3.WeComUser) (toUser []*organization2.User) {

	if fromUser != nil {
		password, _ := organization2.HashPassword(`123456`)
		for _, user := range fromUser {
			toUser = append(toUser, &organization2.User{
				Account:  user.UserId,
				Name:     user.Name,
				NickName: user.Name,
				// todo Position 关联
				DepartmentId:  int64(user.WeComMainDepartmentId),
				MobilePhone:   user.Mobile,
				Gender:        user.Gender,
				Email:         user.Email,
				ExternalEmail: user.Email,
				Avatar:        user.Avatar,
				Password:      password,
				WeComUserId:   user.UserId,
			})
		}
	}

	return toUser
}
