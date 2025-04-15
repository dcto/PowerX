package app

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/request"

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

func (l *CreateWeComAppGroupLogic) CreateWeComAppGroup(opt *types.AppGroupCreateRequest) (resp *types.AppGroupCreateReply, err error) {
	reply, err := l.svcCtx.PowerX.SCRM.WeCom.CreateWeComAppGroupRequest(&request.RequestAppChatCreate{
		Name:     opt.Name,
		Owner:    opt.Owner,
		UserList: opt.UserList,
		ChatID:   opt.ChatId,
	})

	return &types.AppGroupCreateReply{
		ChatId: reply.ChatID,
	}, err
}
