package brainx

import (
	"PowerX/internal/config"
	"PowerX/internal/logic/openapi/provider/brainx/schema"
	providerclient "PowerX/internal/provider"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"time"
)

type BrainXProviderClient struct {
	providerclient.ProviderClientInterface

	AuthToken  *schema.ResponseAuthToken
	BaseURL    string
	httpClient *http.Client
	conf       *config.Config
	cache      cache.CacheInterface
	tokenKey   string
}

func NewBrainXProviderClient(config *config.Config, cache cache.CacheInterface) *BrainXProviderClient {

	return &BrainXProviderClient{
		BaseURL:    config.OpenAPI.Providers.BrainX.BaseUrl,
		httpClient: &http.Client{Timeout: 60 * time.Second},
		conf:       config,
		cache:      cache,
		tokenKey:   "provider.brainx.access_token",
	}
}

func (sp *BrainXProviderClient) Auth(ctx context.Context) (*schema.ResponseAuthToken, error) {
	url := "/auth"

	// 构造请求 body
	body := map[string]string{
		"access_key": sp.conf.OpenAPI.Providers.BrainX.AccessKey,    // 从配置中获取
		"secret_key": sp.conf.OpenAPI.Providers.BrainX.SecretKey,    // 从配置中获取
		"platform":   sp.conf.OpenAPI.Providers.BrainX.ProviderName, // 从配置中获取
	}

	// 发起 POST 请求
	resp, err := sp.HTTPPost(ctx, url, body, false, nil) // use_auth 设置为 false
	if err != nil {
		return nil, err
	}

	// 将 JSON 响应转换为 ResponseAuthToken 对象
	var token *schema.ResponseAuthToken
	err = json.Unmarshal(resp, &token)
	if err != nil {
		return nil, err
	}
	if token.Token.AccessToken == "" {
		return nil, errors.New("auth returned invalid  access token")
	}

	return token, nil
}

func (sp *BrainXProviderClient) GetAccessToken(ctx context.Context) (string, error) {
	// 从缓存中获取 token
	token, err := sp.cache.Get(sp.tokenKey, nil)
	if err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			return "", err
		}
	}

	if token == nil {
		// 如果缓存中没有 token，则调用 Auth 获取新的 token
		sp.AuthToken, err = sp.Auth(ctx)
		if err != nil {
			return "", fmt.Errorf("request powerx provider auth error: %v", err)
		}

		// 将 token 存入缓存，设置过期时间
		expiredIn := sp.AuthToken.Token.ExpiresIn
		err = sp.cache.Set(sp.tokenKey, sp.AuthToken, time.Duration(expiredIn)*time.Second)
		if err != nil {
			return "", err
		}
	} else {
		// 从 map[string]interface{} 转换到 ResponseAuthToken
		tokenMap, ok := token.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("invalid token type in cache")
		}

		// Marshal and Unmarshal to convert
		tokenBytes, err := json.Marshal(tokenMap)
		if err != nil {
			return "", err
		}

		var authToken schema.ResponseAuthToken
		if err := json.Unmarshal(tokenBytes, &authToken); err != nil {
			return "", err
		}
		sp.AuthToken = &authToken
	}

	// 返回 accessToken
	return sp.AuthToken.Token.AccessToken, nil
}
func (sp *BrainXProviderClient) doRequest(ctx context.Context, req *http.Request, useAuth bool) ([]byte, error) {
	// 如果需要授权，添加 Authorization 头部
	if useAuth {
		token, err := sp.GetAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	// 发起 HTTP 请求
	resp, err := sp.httpClient.Do(req)
	logx.Info(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: status code %d", resp.StatusCode)
	}

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (sp *BrainXProviderClient) HTTPGet(ctx context.Context, uri string, params map[string]string, useAuth bool, headers map[string]string) ([]byte, error) {
	url := sp.BaseURL + uri
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// 添加自定义头部信息
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return sp.doRequest(ctx, req, useAuth)
}

func (sp *BrainXProviderClient) HTTPPost(ctx context.Context, uri string, jsonData interface{}, useAuth bool, headers map[string]string) ([]byte, error) {
	// 将 jsonData 转换为 JSON 格式的字节流
	body, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	url := sp.BaseURL + uri
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 添加自定义头部信息
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return sp.doRequest(ctx, req, useAuth)
}

// 从 BrainX 获取 SSE 数据流并将其转发到前端
// 通用的 POST 请求流式传输方法
func (sp *BrainXProviderClient) StreamPOST(
	ctx context.Context, uri string, jsonData interface{},
	w http.ResponseWriter, headers map[string]string,
	commandCallback func(message *schema.SSEMessage) (*schema.SSEMessage, error),
) error {
	// 获取访问令牌
	token, err := sp.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %v", err)
	}

	// 将 jsonData 转换为 JSON 格式的字节流
	body, err := json.Marshal(jsonData)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	// 创建请求
	url := sp.BaseURL + uri
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	// 添加自定义头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := sp.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch SSE stream, status code: %d", resp.StatusCode)
	}

	// 设置响应头为 SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 使用不超时的连接
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("failed to cast response writer to http.Flusher")
	}

	// 逐字流式处理响应数据并即时推送到客户端
	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return errors.New("request timed out")
		default:
			n, err := resp.Body.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("failed to read response body: %v", err)
			}
			if n == 0 {
				continue
			}

			bufMsg := buf[:n]

			// 每次读取到数据立即发送到客户端
			//fmt.Println(string(buf[:n]))
			// 如果需要有回调指令嵌入
			if commandCallback != nil {
				// 先解析当前消息内容
				msg, err := schema.ParseSSEMessage(string(bufMsg))
				if err != nil {
					return fmt.Errorf("failed to parse sse message: %v", err)
				}

				// 如果当前信息已经完成，则嵌入command指令给前端
				if msg.Status == schema.StatusFinished {
					msg, err = commandCallback(msg)
					if err != nil {
						// 如果不影响整体流程，可以继续而不是返回错误
						logx.Errorf("command callback error: %v", err)
						continue
					}

					strMsg, err := schema.BuildSSEMessage(msg)
					if err != nil {
						logx.Errorf("failed to build sse message: %v", err)
						continue
					}
					bufMsg = []byte(strMsg)
					//fmt2.Dump(bufMsg)
				}
			}

			_, err = w.Write(bufMsg)
			if err != nil {
				return fmt.Errorf("failed to write to response: %v", err)
			}

			// 刷新输出缓冲区，确保数据立即推送
			flusher.Flush()

			// 如果已经读取完所有数据，则结束
			if err == io.EOF {
				break
			}
		}
	}

	return nil
}
