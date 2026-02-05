package config

import "encoding/json"

type Config struct {
	GLMAPIKey string `json:"glm_api_key"`
	GLMURL    string `json:"glm_url"`
	Port      string `json:"port"`
}

type PromptConfig struct {
	Version     string   `json:"version"`
	Language    string   `json:"language"`
	Description string   `json:"description"`
	Prompts     []Prompt `json:"prompts"`
}

type Prompt struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Template string          `json:"template"`
	Role     string          `json:"role"`
	Inputs   []PromptInput   `json:"inputs"`
	Output   json.RawMessage `json:"output_schema,omitempty"`
}

type PromptInput struct {
	Key         string `json:"key"`
	Description string `json:"description"`
}
