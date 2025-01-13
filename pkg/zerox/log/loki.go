package log

import (
	fmt2 "PowerX/pkg/printx"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LokiWriter struct {
	logx.Writer
	url        string
	labels     map[string]string
	client     *http.Client
	retryCount int
}

// 新增重试次数参数
func NewLokiWriter(conf LokiConf) *LokiWriter {
	if conf.RetryCount <= 0 {
		conf.RetryCount = 3 // 默认重试3次
	}
	return &LokiWriter{
		url:        conf.URL,
		labels:     conf.Labels,
		client:     &http.Client{Timeout: 10 * time.Second},
		retryCount: conf.RetryCount,
	}
}

// Write 方法实现
func (w *LokiWriter) Write(p []byte) (n int, err error) {
	// 获取当前时间的 Unix 纳秒时间戳
	timestamp := time.Now().UnixNano()

	entry := map[string]interface{}{
		"streams": []map[string]interface{}{
			{
				"stream": w.labels,
				"values": [][]string{
					{
						fmt.Sprintf("%d", timestamp), // 使用 Unix 纳秒时间戳
						string(p),                    // 日志内容
					},
				},
			},
		},
	}

	// 将日志转换为 JSON
	body, err := json.Marshal(entry)
	if err != nil {
		return 0, err
	}
	// fmt2.Dump(string(body))

	// 发送 HTTP 请求到 Loki，带有重试机制
	for attempt := 1; attempt <= w.retryCount; attempt++ {
		resp, err := w.client.Post(w.url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			// 在重试次数内失败时，继续尝试
			if attempt == w.retryCount {
				return 0, fmt.Errorf("failed to send log to Loki: %v", err)
			}
			time.Sleep(time.Second * time.Duration(attempt)) // 按照尝试次数递增等待时间
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			if attempt == w.retryCount {
				errorBody, _ := io.ReadAll(resp.Body) // 读取响应体的内容
				fmt2.Dump(string(errorBody))
				return 0, fmt.Errorf("failed to send log to Loki: status %d - %s, response body: %s", resp.StatusCode, resp.Status, string(errorBody))
			}
			time.Sleep(time.Second * time.Duration(attempt)) // 按照尝试次数递增等待时间
			continue
		}

		return len(p), nil
	}

	return 0, fmt.Errorf("failed to send log to Loki after %d attempts", w.retryCount)
}

// Close 方法
func (w *LokiWriter) Close() error {
	// 可以用于关闭相关资源（如连接池等）
	return nil
}

// 实现 Alert 方法
func (w *LokiWriter) Alert(v any) {
	message := fmt.Sprint(v)
	w.log(message)
}

// 实现 Debug 方法
func (w *LokiWriter) Debug(v any, fields ...logx.LogField) {
	message := fmt.Sprint(v)
	w.logWithFields(message, fields...)
}

// 实现 Error 方法
func (w *LokiWriter) Error(v any, fields ...logx.LogField) {
	message := fmt.Sprint(v)
	w.logWithFields(message, fields...)
}

// 实现 Info 方法
func (w *LokiWriter) Info(v any, fields ...logx.LogField) {
	message := fmt.Sprint(v)
	w.logWithFields(message, fields...)
}

// 实现 Severe 方法
func (w *LokiWriter) Severe(v any) {
	message := fmt.Sprint(v)
	w.log(message)
}

// 实现 Slow 方法
func (w *LokiWriter) Slow(v any, fields ...logx.LogField) {
	message := fmt.Sprint(v)
	w.logWithFields(message, fields...)
}

// 实现 Stack 方法
func (w *LokiWriter) Stack(v any) {
	message := fmt.Sprint(v)
	w.log(message)
}

// 实现 Stat 方法
func (w *LokiWriter) Stat(v any, fields ...logx.LogField) {
	message := fmt.Sprint(v)
	w.logWithFields(message, fields...)
}

// 统一的日志处理函数
func (w *LokiWriter) log(message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		fmt.Printf("failed to send log: %v\n", err)
	}
}

// 处理带有字段的日志
func (w *LokiWriter) logWithFields(message string, fields ...logx.LogField) {
	// 这里可以将字段添加到日志中
	for _, field := range fields {
		message = fmt.Sprintf("%s, %s=%v", message, field.Key, field.Value)
	}
	_, err := w.Write([]byte(message))
	if err != nil {
		fmt.Printf("failed to send log with fields: %v\n", err)
	}
}
