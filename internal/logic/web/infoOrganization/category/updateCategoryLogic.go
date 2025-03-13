package category

import (
	category2 "PowerX/internal/logic/admin/infoOrganization/category"
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/model/powermodel"
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryRequest) (resp *types.UpdateCategoryReply, err error) {
	vAuthCustomer := l.ctx.Value(customerDomain.AuthCustomerKey)
	authCustomer := vAuthCustomer.(*customerDomain2.Customer)

	existedItem, err := l.svcCtx.PowerX.Category.CategoryRepository.GetByID(l.ctx, req.Id, nil)
	if err != nil {
		return nil, err
	}

	if existedItem == nil {
		return nil, errorx.ErrNotFoundObject
	}

	if existedItem.CustomerId != authCustomer.Id {
		return nil, errorx.ErrCustomerNotMatch
	}

	category := category2.TransformRequestToCategory(&req.Category)
	category.PowerModel = powermodel.PowerModel{
		Id: req.Id,
	}
	category.CustomerId = authCustomer.Id

	category, err = l.svcCtx.PowerX.Category.UpsertCategory(l.ctx, category)

	if err != nil {
		return nil, err
	}

	return &types.UpdateCategoryReply{
		Id: category.Id,
	}, nil

}
