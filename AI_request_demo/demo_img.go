package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetImageBase64 将图片转换为Base64编码
func GetImageBase64(path string) (string, string, error) {
	var imgByte []byte
	var _ error

	if strings.Contains(path, "http") {
		// 获取网络图片
		client := &http.Client{
			Timeout: time.Second * 5,
		}

		resp, err := client.Get(path)
		if err != nil {
			return "", "", fmt.Errorf("获取网络图片失败: %v", err)
		}
		defer resp.Body.Close()

		imgByte, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", "", fmt.Errorf("读取网络图片数据失败: %v", err)
		}
	} else {
		// 获取本地文件
		file, err := os.Open(path)
		if err != nil {
			return "", "", fmt.Errorf("获取本地图片失败: %v", err)
		}
		defer file.Close()

		imgByte, err = io.ReadAll(file)
		if err != nil {
			return "", "", fmt.Errorf("读取本地图片数据失败: %v", err)
		}
	}

	// 检测图片类型
	mimeType := http.DetectContentType(imgByte)
	base64String := base64.StdEncoding.EncodeToString(imgByte)
	return base64String, mimeType, nil
}

// GLM多模态API请求结构体
type GLMChatRequest struct {
	Model     string       `json:"model"`
	Messages  []GLMMessage `json:"messages"`
	MaxTokens int          `json:"max_tokens,omitempty"`
}

type GLMMessage struct {
	Role    string       `json:"role"`
	Content []GLMContent `json:"content"`
}

type GLMContent struct {
	Type     string       `json:"type"`
	Text     string       `json:"text,omitempty"`
	ImageURL *GLMImageURL `json:"image_url,omitempty"`
}

type GLMImageURL struct {
	URL string `json:"url"`
}

// GLM多模态API响应结构体
type GLMChatResponse struct {
	Choices []GLMChoice `json:"choices"`
}

type GLMChoice struct {
	Message GLMResponseMessage `json:"message"`
}

type GLMResponseMessage struct {
	Content string `json:"content"`
}

// 调用GLM多模态模型进行图像分析
func analyzeImageWithGLM(imageBase64, imageType, question string) error {
	apiKey := "0ffcfbc67b30483a93a6fa041936f76b.OS41RBK74yt6mskO" // 替换为你的API Key
	url := "https://open.bigmodel.cn/api/paas/v4/chat/completions"

	// 构建多模态消息
	messages := []GLMMessage{
		{
			Role: "user",
			Content: []GLMContent{
				{
					Type: "text",
					Text: question,
				},
				{
					Type: "image_url",
					ImageURL: &GLMImageURL{
						URL: fmt.Sprintf("data:%s;base64,%s", imageType, imageBase64),
					},
				},
			},
		},
	}

	requestData := GLMChatRequest{
		Model:     "glm-4v",
		Messages:  messages,
		MaxTokens: 2048,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("JSON序列化失败: %v", err)
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("API请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析响应
	var response GLMChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("JSON解析失败: %v", err)
	}

	if len(response.Choices) > 0 {
		fmt.Printf("分析结果: %s\n", response.Choices[0].Message.Content)
	} else {
		return fmt.Errorf("未收到有效回复")
	}

	return nil
}

func main() {
	// 示例用法
	fmt.Println("GLM多模态图片分析调试工具")
	fmt.Println(strings.Repeat("=", 50))

	// 方式1: 本地图片路径
	imagePath := "./img.png" // 替换为你的图片路径

	// 方式2: 网络图片URL
	// imagePath := "https://example.com/image.jpg"

	question := "请详细描述这张图片的内容"

	// 将图片转换为Base64
	imageBase64, imageType, err := GetImageBase64(imagePath)
	if err != nil {
		log.Fatalf("图片处理失败: %v", err)
	}

	fmt.Printf("图片类型: %s\n", imageType)
	fmt.Printf("Base64数据长度: %d 字符\n", len(imageBase64))
	fmt.Println(strings.Repeat("=", 50))

	// 调用GLM模型分析图片
	err = analyzeImageWithGLM(imageBase64, imageType, question)
	if err != nil {
		log.Fatalf("分析失败: %v", err)
	}
}
