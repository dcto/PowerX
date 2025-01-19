package ai

import (
	"PowerX/internal/config"
	"PowerX/internal/logic/openapi/provider/brainx/schema"
	"PowerX/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type ChatBotUseCase struct {
	db *gorm.DB
}

func NewChatBotUseCase(conf *config.Config, db *gorm.DB) (uc *ChatBotUseCase) {

	uc = &ChatBotUseCase{
		db: db,
	}

	return uc
}

func (uc *ChatBotUseCase) Chat(ctx context.Context, req *types.ChatRequest, w http.ResponseWriter) error {
	msg := "这是一个对话模式"
	err := uc.StreamingResponse(ctx, msg, w)
	return err
}

func (uc *ChatBotUseCase) AgentChat(ctx context.Context, req *types.ChatRequest, w http.ResponseWriter) error {
	msg := "这是一个代理对话模式"
	err := uc.StreamingResponse(ctx, msg, w)
	return err
}

func (uc *ChatBotUseCase) StreamingResponse(ctx context.Context, responseMessage string, w http.ResponseWriter) error {
	// 设置响应头为 SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 使用不超时的连接
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("failed to cast response writer to http.Flusher")
	}

	// 初始化消息并发送状态：processing
	initialMessage := schema.SSEMessage{
		Status: "processing", // 初始状态
	}

	initialMessageJSON, err := json.Marshal(initialMessage)
	if err != nil {
		return err
	}

	// 发送初始状态的消息
	_, err = w.Write([]byte(fmt.Sprintf("data: %s\n\n", string(initialMessageJSON))))
	if err != nil {
		return err
	}
	flusher.Flush()

	// 模拟外部数据源的响应（例如调用一个API或者数据库查询）

	// 模拟逐字返回 JSON 格式的 SSE 数据
	for _, char := range responseMessage {
		select {
		case <-ctx.Done():
			// 如果上下文超时，退出
			return errors.New("request timed out")
		default:
			// 每次发送一个字母
			sseMessage := schema.SSEMessage{
				Status:  "data",       // 传输中的状态
				Content: string(char), // 当前的字符内容
				//Message: fmt.Sprintf("正在发送字符: %c", char), // 当前字符的信息
			}

			// 将 SSEMessage 转换为 JSON 格式
			messageJSON, err := json.Marshal(sseMessage)
			if err != nil {
				// 如果转换失败，返回错误信息
				return err
			}

			// 发送消息到客户端
			messageStr := fmt.Sprintf("data: %s\n\n", string(messageJSON))
			_, err = w.Write([]byte(messageStr))
			if err != nil {
				return err
			}

			// 刷新输出缓冲区
			flusher.Flush()

			// 模拟逐字延时
			time.Sleep(20 * time.Millisecond) // 每个字符发送延时20毫秒
		}
	}

	// 完成所有消息后，发送结束标志
	doneMessage := schema.SSEMessage{
		Status:  "finished", // 完成状态
		Message: "所有消息已发送完毕。",
	}

	doneJSON, err := json.Marshal(doneMessage)
	if err != nil {
		return err
	}

	// 发送完成消息
	_, err = w.Write([]byte(fmt.Sprintf("data: %s\n\n", string(doneJSON))))
	if err != nil {
		return err
	}

	// 刷新输出缓冲区
	flusher.Flush()

	return nil
}
