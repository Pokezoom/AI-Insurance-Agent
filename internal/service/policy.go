package service

import (
	"AI-Insurance-Agent/config"
	"AI-Insurance-Agent/internal/client"
	"AI-Insurance-Agent/internal/model"
	"AI-Insurance-Agent/internal/repository"
	"encoding/json"
	"os"
	"strings"
)

type PolicyService struct {
	glmClient    *client.GLMClient
	promptConfig *config.PromptConfig
	recordRepo   *repository.RecordRepository
}

func NewPolicyService(glmClient *client.GLMClient, recordRepo *repository.RecordRepository) (*PolicyService, error) {
	data, err := os.ReadFile("prompt/prompt.json")
	if err != nil {
		return nil, err
	}

	var promptConfig config.PromptConfig
	json.Unmarshal(data, &promptConfig)

	return &PolicyService{
		glmClient:    glmClient,
		promptConfig: &promptConfig,
		recordRepo:   recordRepo,
	}, nil
}

func (s *PolicyService) AnalyzePolicy(userID int64, imageBase64, imageType string) (*model.AnalysisRecord, error) {
	record := &model.AnalysisRecord{
		UserID:    userID,
		ImageType: imageType,
		Status:    "pending",
	}
	s.recordRepo.Create(record)

	prompt := s.getPromptByID("policy_structuring")
	if prompt == nil {
		record.Status = "failed"
		record.ErrorMessage = "未找到prompt配置"
		s.recordRepo.Update(record)
		return record, nil
	}

	finalPrompt := strings.ReplaceAll(prompt.Template, "{{policy_text}}", "请分析这张保单图片")
	result, err := s.glmClient.AnalyzeImage(imageBase64, imageType, finalPrompt)

	if err != nil {
		record.Status = "failed"
		record.ErrorMessage = err.Error()
	} else {
		record.Status = "success"
		record.StructuredData = result
		record.AnalysisResult = result
	}

	s.recordRepo.Update(record)
	return record, err
}

func (s *PolicyService) GetRecord(recordID, userID int64) (*model.AnalysisRecord, error) {
	return s.recordRepo.FindByID(recordID)
}

func (s *PolicyService) ListRecords(userID int64, page, pageSize int) ([]model.AnalysisRecord, int64, error) {
	return s.recordRepo.ListByUserID(userID, page, pageSize)
}

func (s *PolicyService) DeleteRecord(recordID, userID int64) error {
	return s.recordRepo.Delete(recordID, userID)
}

func (s *PolicyService) getPromptByID(id string) *config.Prompt {
	for _, p := range s.promptConfig.Prompts {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
