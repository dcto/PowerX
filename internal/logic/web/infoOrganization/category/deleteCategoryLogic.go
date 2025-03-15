package category

import (
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryRequest) (resp *types.DeleteCategoryReply, err error) {
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

	_, err = l.svcCtx.PowerX.Category.CategoryRepository.Delete(l.ctx, map[string]interface{}{
		"id": req.Id,
	}, existedItem, false)

	if err != nil {
		return nil, err
	}

	return &types.DeleteCategoryReply{
		Id: existedItem.Id,
	}, nil
}
