package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ChatCompletionRequest 定义请求结构体，匹配API要求的格式
type ChatCompletionRequest struct {
	Model     string                  `json:"model"`
	Messages  []ChatCompletionMessage `json:"messages"`
	MaxTokens int                     `json:"max_tokens,omitempty"`
}
type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionResponse 定义响应结构体，用于解析返回结果
type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := "0ffcfbc67b30483a93a6fa041936f76b.OS41RBK74yt6mskO"
	url := "https://open.bigmodel.cn/api/paas/v4/chat/completions"

	// 1. 构建请求数据
	requestData := ChatCompletionRequest{
		Model: "glm-4",
		Messages: []ChatCompletionMessage{
			{Role: "user", Content: "你好，请介绍一下你自己。"},
		},
		MaxTokens: 1024,
	}
	jsonData, _ := json.Marshal(requestData)

	// 2. 创建HTTP请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey) // 设置认证头

	// 3. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// 4. 解析响应
	var response ChatCompletionResponse
	json.Unmarshal(body, &response)
	if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	} else {
		fmt.Println("未收到有效回复。", resp)
	}
}
