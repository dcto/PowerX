package category

import (
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"PowerX/pkg/mapx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchCategoryLogic {
	return &PatchCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchCategoryLogic) PatchCategory(req *types.PatchCategoryRequest) (resp *types.PatchCategoryReply, err error) {
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
	updateValues := mapx.MapUpdatesFromObject(req.Category)

	existedItem, err = l.svcCtx.PowerX.Category.CategoryRepository.Patch(l.ctx, map[string]interface{}{
		"id": req.Id,
	}, updateValues)

	if err != nil {
		return nil, err
	}

	return &types.PatchCategoryReply{
		TransformCategoryToReplyForWeb(existedItem),
	}, nil
}
