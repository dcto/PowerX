package organization

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeComDepartmentTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询组织架构
func NewGetWeComDepartmentTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeComDepartmentTreeLogic {
	return &GetWeComDepartmentTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWeComDepartmentTreeLogic) GetWeComDepartmentTree(req *types.GetWecomDepartmentTreeRequest) (resp *types.GetWecomDepartmentTreeReply, err error) {
	res, err := l.svcCtx.PowerX.SCRM.WeCom.GetDepartmentBy(l.ctx, req.DepartmentId, 5)
	if err != nil {
		return nil, err
	}

	return &types.GetWecomDepartmentTreeReply{
		Department: TransformWeComDepartmentToReply(res),
	}, nil
}
