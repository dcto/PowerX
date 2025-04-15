package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWeComAppGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// App创建企业群
func NewCreateWeComAppGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWeComAppGroupLogic {
	return &CreateWeComAppGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWeComAppGroupLogic) CreateWeComAppGroup(req *types.AppGroupCreateRequest) (resp *types.AppGroupCreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
