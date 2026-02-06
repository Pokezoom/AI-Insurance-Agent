package model

import "time"

type AnalysisRecord struct {
	ID             int64     `gorm:"primaryKey" json:"id"`
	UserID         int64     `gorm:"not null;index" json:"user_id"`
	ImageURL       string    `json:"image_url"`
	ImageType      string    `json:"image_type"`
	StructuredData string    `gorm:"type:json" json:"structured_data"`
	AnalysisResult string    `gorm:"type:text" json:"analysis_result"`
	Status         string    `gorm:"default:pending" json:"status"`
	ErrorMessage   string    `gorm:"type:text" json:"error_message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
