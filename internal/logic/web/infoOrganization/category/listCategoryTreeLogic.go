package category

import (
	customerDomain2 "PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/uc/powerx/crm/customerDomain"
	"PowerX/internal/uc/powerx/crm/infoOrganization"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryTreeLogic {
	return &ListCategoryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoryTreeLogic) ListCategoryTree(req *types.ListCategoryTreeRequest) (resp *types.ListCategoryTreeReply, err error) {
	vAuthCustomer := l.ctx.Value(customerDomain.AuthCustomerKey)
	authCustomer := vAuthCustomer.(*customerDomain2.Customer)

	option := infoOrganization.FindCategoryOption{
		Names:      req.Names,
		OrderBy:    req.OrderBy,
		CustomerId: authCustomer.Id,
	}

	// 获取模型类型的列表
	productCategoryTree := l.svcCtx.PowerX.Category.ListCategoryTree(l.ctx, &option, 0)

	// 转化返回类型的列表
	productCategoryReplyList := TransformCategoriesToReplyForWeb(productCategoryTree)

	return &types.ListCategoryTreeReply{
		ProductCategories: productCategoryReplyList,
	}, nil
}
