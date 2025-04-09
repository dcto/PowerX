package customerDomain

import (
	"PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/model/powerModel"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"PowerX/pkg/securityx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type CustomerUseCase struct {
	db *gorm.DB
}

func NewCustomerUseCase(db *gorm.DB) *CustomerUseCase {
	return &CustomerUseCase{
		db: db,
	}
}

type FindManyCustomersOption struct {
	LikeName   string
	LikeMobile string
	Mobile     string
	Statuses   []int
	Sources    []int
	OrderBy    string
	types.PageEmbedOption
}

func (uc *CustomerUseCase) buildFindQueryNoPage(db *gorm.DB, opt *FindManyCustomersOption) *gorm.DB {
	if opt.LikeName != "" {
		db = db.Where("name LIKE ?", "%"+opt.LikeName+"%")
	}
	if opt.LikeMobile != "" {
		db = db.Where("mobile LIKE ?", "%"+opt.LikeMobile+"%")
	}
	if opt.Mobile != "" {
		db = db.Where("mobile = ?", opt.Mobile)
	}
	if len(opt.Statuses) > 0 {
		db = db.Where("status IN ?", opt.Statuses)
	}
	if len(opt.Sources) > 0 {
		db = db.Where("source IN ?", opt.Sources)
	}
	orderBy := "id desc"
	if opt.OrderBy != "" {
		orderBy = opt.OrderBy + "," + orderBy
	}
	db.Order(orderBy)

	return db
}

func (uc *CustomerUseCase) FindManyCustomers(ctx context.Context, opt *FindManyCustomersOption) (pageList types.Page[*customerDomain.Customer], err error) {
	var customers []*customerDomain.Customer
	db := uc.db.WithContext(ctx).Model(&customerDomain.Customer{})

	db = uc.buildFindQueryNoPage(db, opt)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		panic(err)
	}

	opt.DefaultPageIfNotSet()
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		db.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	if err := db.
		//Debug().
		Find(&customers).Error; err != nil {
		panic(err)
	}

	return types.Page[*customerDomain.Customer]{
		List:      customers,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, nil
}

func (uc *CustomerUseCase) CreateCustomer(ctx context.Context, customer *customerDomain.Customer) error {
	if err := uc.db.WithContext(ctx).Create(&customer).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrDuplicatedInsert, "该对象不能重复创建")
		}
		panic(err)
	}
	return nil
}

func (uc *CustomerUseCase) CreateCustomerByRegisterCode(ctx context.Context, customer *customerDomain.Customer, registerCode *customerDomain.RegisterCode) error {

	err := uc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&customerDomain.Customer{}).Create(&customer).Error
		if err != nil {
			return err
		}

		// 更新注册记录的受邀者的ID
		registerCode.RegisterCustomerID = customer.Id
		err = tx.Model(registerCode).
			Update("register_customer_id", customer.Id).Error
		return err
	})

	return err
}

func (uc *CustomerUseCase) UpsertCustomer(ctx context.Context, customer *customerDomain.Customer) (*customerDomain.Customer, error) {

	customers := []*customerDomain.Customer{customer}

	err := uc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := powerModel.UpsertModelsOnUniqueID(tx, &customerDomain.Customer{}, customerDomain.CustomerUniqueId, customers, nil, false)

		if err != nil {
			panic(errors.Wrap(err, "upsert customerDomain failed"))
		}
		// 如果是新增用户，那么需要给一个唯一识别号
		if customer.Uuid == "" {
			customer.Uuid = securityx.GenerateUUIDString()
			err = powerModel.UpsertModelsOnUniqueID(tx, &customerDomain.Customer{}, customerDomain.CustomerUniqueId, customer, []string{"uuid"}, false)
			if err != nil {
				return err
			}
		}
		return err
	})

	return customer, err
}

func (uc *CustomerUseCase) UpsertCustomers(ctx context.Context, customers []*customerDomain.Customer) ([]*customerDomain.Customer, error) {

	err := powerModel.UpsertModelsOnUniqueID(uc.db.WithContext(ctx), &customerDomain.Customer{}, customerDomain.CustomerUniqueId, customers, nil, false)

	if err != nil {
		panic(errors.Wrap(err, "batch upsert customers failed"))
	}

	return customers, err
}

func (uc *CustomerUseCase) UpdateCustomer(ctx context.Context, id int64, customer *customerDomain.Customer) error {
	//fmt.Dump(customer)
	err := uc.db.WithContext(ctx).Model(&customerDomain.Customer{}).
		//Debug().
		Where(id).Updates(&customer).Error

	return err
}

func (uc *CustomerUseCase) GetCustomer(ctx context.Context, id int64) (*customerDomain.Customer, error) {
	var customer customerDomain.Customer
	if err := uc.db.WithContext(ctx).First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到客户")
		}
		panic(err)
	}
	return &customer, nil
}

func (uc *CustomerUseCase) GetCustomerByMobile(ctx context.Context, mobile string) (*customerDomain.Customer, error) {
	var customer customerDomain.Customer
	if err := uc.db.WithContext(ctx).
		Where("mobile", mobile).
		First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到客户")
		}
		panic(err)
	}
	return &customer, nil
}

func (uc *CustomerUseCase) GetCustomerByUUID(ctx context.Context, uuid string) (*customerDomain.Customer, error) {
	var customer customerDomain.Customer
	if err := uc.db.WithContext(ctx).
		Where("uuid", uuid).
		First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到客户")
		}
		panic(err)
	}
	return &customer, nil
}
func (uc *CustomerUseCase) GetCustomerByInviteCode(ctx context.Context, inviteCode string) (*customerDomain.Customer, error) {
	var customer customerDomain.Customer
	if err := uc.db.WithContext(ctx).
		Where("invite_code", inviteCode).
		First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到客户")
		}
		panic(err)
	}
	return &customer, nil
}

func (uc *CustomerUseCase) DeleteCustomer(ctx context.Context, id int64) error {
	result := uc.db.WithContext(ctx).Delete(&customerDomain.Customer{}, id)
	if err := result.Error; err != nil {
		panic(err)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "未找到客户")
	}
	return nil
}

func (uc *CustomerUseCase) CheckRegisterPhoneExist(ctx context.Context, mobile string) bool {

	customer := &customerDomain.Customer{}
	err := uc.db.WithContext(ctx).
		//Debug().
		Unscoped().
		Where("mobile", mobile).
		First(customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		panic(err)
	}

	return customer.Id > 0

}
