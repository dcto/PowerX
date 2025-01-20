package schema

import (
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/pkg/errors"
	"strings"
)

type SSEStatus string

const (
	StatusProcessing SSEStatus = "processing"
	StatusData       SSEStatus = "data"
	StatusFinished   SSEStatus = "finished"
	StatusError      SSEStatus = "error"
)

type Command struct {
	Type    string         `json:"type"`
	Payload object.HashMap `json:"payload"`
}

// SSEMessage 定义 SSE 返回的数据格式
type SSEMessage struct {
	Status  SSEStatus `json:"status"`
	Content string    `json:"content,omitempty"`
	Message string    `json:"message,omitempty"`
	Command Command   `json:"command,omitempty"`
}

// 解析 data 字符串为 SSEMessage
func ParseSSEMessage(data string) (*SSEMessage, error) {
	// 去掉 `data: ` 前缀
	if !strings.HasPrefix(data, "data: ") {
		return nil, errors.New("invalid data format")
	}
	rawJSON := strings.TrimPrefix(data, "data: ")

	// 解析 JSON
	var msg SSEMessage
	err := json.Unmarshal([]byte(rawJSON), &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SSE message: %v", err)
	}

	return &msg, nil
}

func BuildSSEMessage(msg *SSEMessage) (string, error) {
	if msg == nil {
		return "", fmt.Errorf("SSEMessage cannot be nil")
	}

	// 将 SSEMessage 编码为 JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to encode SSEMessage: %v", err)
	}

	// 添加 `data: ` 前缀
	return fmt.Sprintf("data: %s\n\n", string(jsonData)), nil
}
