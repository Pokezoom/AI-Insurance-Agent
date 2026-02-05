package service

import (
	"AI-Insurance-Agent/config"
	"AI-Insurance-Agent/internal/client"
	"encoding/json"
	"os"
	"strings"
)

type PolicyService struct {
	glmClient    *client.GLMClient
	promptConfig *config.PromptConfig
}

func NewPolicyService(glmClient *client.GLMClient) (*PolicyService, error) {
	data, err := os.ReadFile("prompt.json")
	if err != nil {
		return nil, err
	}

	var promptConfig config.PromptConfig
	json.Unmarshal(data, &promptConfig)

	return &PolicyService{
		glmClient:    glmClient,
		promptConfig: &promptConfig,
	}, nil
}

func (s *PolicyService) AnalyzePolicy(imageBase64, imageType string) (string, error) {
	prompt := s.getPromptByID("policy_structuring")
	if prompt == nil {
		return "", nil
	}

	finalPrompt := strings.ReplaceAll(prompt.Template, "{{policy_text}}", "请分析这张保单图片")
	return s.glmClient.AnalyzeImage(imageBase64, imageType, finalPrompt)
}

func (s *PolicyService) getPromptByID(id string) *config.Prompt {
	for _, p := range s.promptConfig.Prompts {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
