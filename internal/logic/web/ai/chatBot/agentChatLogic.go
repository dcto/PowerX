package chatBot

import (
	"PowerX/internal/provider/brainx"
	"context"
	"net/http"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgentChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI ChatBot Agent对话
func NewAgentChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgentChatLogic {
	return &AgentChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgentChatLogic) AgentChat(req *types.ChatRequest, w http.ResponseWriter) error {
	//err := l.svcCtx.PowerX.ChatBot.AgentChat(l.ctx, req, w)

	// 向 BrainX 发送请求获取 SSE 数据流
	brainXService := brainx.NewBrainXServiceProvider(l.svcCtx)
	err := brainXService.AgentChat(l.ctx, req, w)

	return err
}
