package category

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

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

	// 获取模型类型的列表
	productCategoryTree, err := l.svcCtx.PowerX.Category.GetCategory(l.ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}

	// 转化返回类别
	categoryReply := TransformCategoryToReplyForWeb(productCategoryTree)

	return &types.GetCategoryReply{
		Category: categoryReply,
	}, nil
}
