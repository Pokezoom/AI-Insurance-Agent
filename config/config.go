package config

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	JWT struct {
		Secret      string `yaml:"secret"`
		ExpireHours int    `yaml:"expire_hours"`
	} `yaml:"jwt"`
	GLM struct {
		APIKey string `yaml:"api_key"`
		URL    string `yaml:"url"`
	} `yaml:"glm"`
}

func LoadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
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
