package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GLMClient struct {
	APIKey string
	URL    string
}

type GLMRequest struct {
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

type GLMResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *GLMClient) AnalyzeImage(imageBase64, imageType, prompt string) (string, error) {
	messages := []GLMMessage{
		{
			Role: "user",
			Content: []GLMContent{
				{Type: "text", Text: prompt},
				{Type: "image_url", ImageURL: &GLMImageURL{
					URL: fmt.Sprintf("data:%s;base64,%s", imageType, imageBase64),
				}},
			},
		},
	}

	reqData := GLMRequest{
		Model:     "glm-4v",
		Messages:  messages,
		MaxTokens: 2048,
	}

	jsonData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response GLMResponse
	json.Unmarshal(body, &response)

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("未收到有效回复")
}
