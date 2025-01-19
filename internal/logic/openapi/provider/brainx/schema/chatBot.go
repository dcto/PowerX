package schema

// SSEMessage 定义 SSE 返回的数据格式
type SSEMessage struct {
	Status  string `json:"status"`
	Content string `json:"content,omitempty"`
	Message string `json:"message,omitempty"`
}
