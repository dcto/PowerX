package customer

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendWeComCustomerGroupMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户群发信息
func NewSendWeComCustomerGroupMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendWeComCustomerGroupMessageLogic {
	return &SendWeComCustomerGroupMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendWeComCustomerGroupMessageLogic) SendWeComCustomerGroupMessage(req *types.WeComAddMsgTemplateRequest) (resp *types.WeComAddMsgTemplateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
