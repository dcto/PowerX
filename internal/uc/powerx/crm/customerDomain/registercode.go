package customerdomain

import (
	"PowerX/internal/model/crm/customerdomain"
	"PowerX/internal/model/powermodel"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type RegisterCodeUseCase struct {
	db *gorm.DB
}

func NewRegisterCodeUseCase(db *gorm.DB) *RegisterCodeUseCase {
	return &RegisterCodeUseCase{
		db: db,
	}
}

type FindManyRegisterCodesOption struct {
	LikeName   string
	LikeMobile string
	Statuses   []int
	Sources    []int
	OrderBy    string
	types.PageEmbedOption
}

func (uc *RegisterCodeUseCase) buildFindQueryNoPage(db *gorm.DB, opt *FindManyRegisterCodesOption) *gorm.DB {
	if opt.LikeName != "" {
		db = db.Where("name LIKE ?", "%"+opt.LikeName+"%")
	}
	if opt.LikeMobile != "" {
		db = db.Where("mobile LIKE ?", "%"+opt.LikeMobile+"%")
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

func (uc *RegisterCodeUseCase) FindManyRegisterCodes(ctx context.Context, opt *FindManyRegisterCodesOption) (pageList types.Page[*customerDomain.RegisterCode], err error) {
	var registerCodes []*customerDomain.RegisterCode
	db := uc.db.WithContext(ctx).Model(&customerDomain.RegisterCode{})

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
		Find(&registerCodes).Error; err != nil {
		panic(err)
	}

	return types.Page[*customerDomain.RegisterCode]{
		List:      registerCodes,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, nil
}

func (uc *RegisterCodeUseCase) CreateRegisterCode(ctx context.Context, registerCode *customerDomain.RegisterCode) error {
	if err := uc.db.WithContext(ctx).Create(&registerCode).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrDuplicatedInsert, "该对象不能重复创建")
		}
		panic(err)
	}
	return nil
}

func (uc *RegisterCodeUseCase) UpsertRegisterCode(ctx context.Context, registerCode *customerDomain.RegisterCode) (*customerDomain.RegisterCode, error) {

	registerCodes := []*customerDomain.RegisterCode{registerCode}

	_, err := uc.UpsertRegisterCodes(ctx, registerCodes)
	if err != nil {
		panic(errors.Wrap(err, "upsert registerCode failed"))
	}

	return registerCode, err
}

func (uc *RegisterCodeUseCase) UpsertRegisterCodes(ctx context.Context, registerCodes []*customerDomain.RegisterCode) ([]*customerDomain.RegisterCode, error) {

	err := powermodel.UpsertModelsOnUniqueID(uc.db.WithContext(ctx), &customerDomain.RegisterCode{}, customerDomain.RegisterCodeUniqueId, registerCodes, nil, false)

	if err != nil {
		panic(errors.Wrap(err, "batch upsert registerCodes failed"))
	}

	return registerCodes, err
}

func (uc *RegisterCodeUseCase) UpdateRegisterCode(ctx context.Context, id int64, registerCode *customerDomain.RegisterCode) error {
	if err := uc.db.WithContext(ctx).Model(&customerDomain.RegisterCode{}).
		//Debug().
		Where(id).Updates(&registerCode).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrDuplicatedInsert, "该对象不能重复创建")
		}
		panic(err)
	}
	return nil
}

func (uc *RegisterCodeUseCase) GetRegisterCodeByCode(ctx context.Context, code string) (*customerDomain.RegisterCode, error) {
	var registerCode customerDomain.RegisterCode
	if err := uc.db.WithContext(ctx).
		//Debug().
		Where("code", code).
		Where("register_customer_id=0").
		First(&registerCode).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到注册码")
		}
		panic(err)
	}
	return &registerCode, nil
}

func (uc *RegisterCodeUseCase) GetRegisterCode(ctx context.Context, id int64) (*customerDomain.RegisterCode, error) {
	var registerCode customerDomain.RegisterCode
	if err := uc.db.WithContext(ctx).First(&registerCode, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到注册码")
		}
		panic(err)
	}
	return &registerCode, nil
}

func (uc *RegisterCodeUseCase) DeleteRegisterCode(ctx context.Context, id int64) error {
	result := uc.db.WithContext(ctx).Delete(&customerDomain.RegisterCode{}, id)
	if err := result.Error; err != nil {
		panic(err)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "未找到注册码")
	}
	return nil
}
