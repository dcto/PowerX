package category

import (
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryRequest) (resp *types.CreateCategoryReply, err error) {
	vAuthCustomer := l.ctx.Value(customerDomain.AuthCustomerKey)
	authCustomer := vAuthCustomer.(*customerDomain2.Customer)

	// Web端创建用户的客户分类
	category := TransformCategoryRequestToCategory(req)
	category.CustomerId = authCustomer.Id
	res, err := l.svcCtx.PowerX.Category.CategoryRepository.Create(l.ctx, category)
	if err != nil {
		return nil, err
	}
	return &types.CreateCategoryReply{
		Category: TransformCategoryToReplyForWeb(res),
	}, nil
}
