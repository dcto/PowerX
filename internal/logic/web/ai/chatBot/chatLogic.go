package chatBot

import (
	"PowerX/internal/provider/brainx"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewChatLogic 创建新的聊天逻辑
func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 模拟从外部服务获取数据流并逐步推送
func (l *ChatLogic) Chat(req *types.ChatRequest, w http.ResponseWriter) error {
	// 向 BrainX 发送请求获取 SSE 数据流
	brainXService := brainx.NewBrainXServiceProvider(l.svcCtx)
	err := brainXService.Chat(l.ctx, req, w)

	return err
}
