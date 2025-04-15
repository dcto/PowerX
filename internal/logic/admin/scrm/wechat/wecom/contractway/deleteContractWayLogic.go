package contractway

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContractWayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除活码
func NewDeleteContractWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContractWayLogic {
	return &DeleteContractWayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContractWayLogic) DeleteContractWay(req *types.DeleteContractWayRequest) (resp *types.DeleteContractWayReply, err error) {
	// todo: add your logic here and delete this line

	return
}
