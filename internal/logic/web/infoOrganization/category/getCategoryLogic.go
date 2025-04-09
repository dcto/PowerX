package category

import (
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.GetCategoryRequest) (resp *types.GetCategoryReply, err error) {
	vAuthCustomer := l.ctx.Value(customerDomain.AuthCustomerKey)
	authCustomer := vAuthCustomer.(*customerDomain2.Customer)

	// 获取模型类型的列表
	categoryTree, err := l.svcCtx.PowerX.Category.GetCategory(l.ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}

	if categoryTree.CustomerId != authCustomer.Id {
		return nil, errorx.ErrCustomerNotMatch
	}

	// 转化返回类别
	categoryReply := TransformCategoryToReplyForWeb(categoryTree)

	return &types.GetCategoryReply{
		Category: categoryReply,
	}, nil
}
