package customerDomain

import (
	"PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/model/powermodel"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type LeadUseCase struct {
	db *gorm.DB
}

func NewLeadUseCase(db *gorm.DB) *LeadUseCase {
	return &LeadUseCase{
		db: db,
	}
}

type FindManyLeadsOption struct {
	LikeName   string
	LikeMobile string
	Statuses   []int
	Sources    []int
	OrderBy    string
	types.PageEmbedOption
}

func (uc *LeadUseCase) buildFindQueryNoPage(db *gorm.DB, opt *FindManyLeadsOption) *gorm.DB {
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

func (uc *LeadUseCase) FindManyLeads(ctx context.Context, opt *FindManyLeadsOption) (pageList types.Page[*customerDomain.Lead], err error) {
	var leads []*customerDomain.Lead
	db := uc.db.WithContext(ctx).Model(&customerDomain.Lead{})

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
		Find(&leads).Error; err != nil {
		panic(err)
	}

	return types.Page[*customerDomain.Lead]{
		List:      leads,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, nil
}

func (uc *LeadUseCase) CreateLead(ctx context.Context, lead *customerDomain.Lead) error {
	if err := uc.db.WithContext(ctx).Create(&lead).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrDuplicatedInsert, "该对象不能重复创建")
		}
		panic(err)
	}
	return nil
}

func (uc *LeadUseCase) UpsertLead(ctx context.Context, lead *customerDomain.Lead) (*customerDomain.Lead, error) {

	leads := []*customerDomain.Lead{lead}

	_, err := uc.UpsertLeads(ctx, leads)
	if err != nil {
		panic(errors.Wrap(err, "upsert lead failed"))
	}

	return lead, err
}

func (uc *LeadUseCase) UpsertLeads(ctx context.Context, leads []*customerDomain.Lead) ([]*customerDomain.Lead, error) {

	err := powerModel.UpsertModelsOnUniqueID(uc.db.WithContext(ctx), &customerDomain.Lead{}, customerDomain.LeadUniqueId, leads, nil, false)

	if err != nil {
		panic(errors.Wrap(err, "batch upsert leads failed"))
	}

	return leads, err
}

func (uc *LeadUseCase) UpdateLead(ctx context.Context, id int64, lead *customerDomain.Lead) error {
	if err := uc.db.WithContext(ctx).Model(&customerDomain.Lead{}).
		//Debug().
		Where(id).Updates(&lead).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrDuplicatedInsert, "该对象不能重复创建")
		}
		panic(err)
	}
	return nil
}

func (uc *LeadUseCase) GetLead(ctx context.Context, id int64) (*customerDomain.Lead, error) {
	var lead customerDomain.Lead
	if err := uc.db.WithContext(ctx).First(&lead, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "未找到线索")
		}
		panic(err)
	}
	return &lead, nil
}

func (uc *LeadUseCase) DeleteLead(ctx context.Context, id int64) error {
	result := uc.db.WithContext(ctx).Delete(&customerDomain.Lead{}, id)
	if err := result.Error; err != nil {
		panic(err)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "未找到线索")
	}
	return nil
}
