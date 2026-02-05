package model

type PolicyAnalysisRequest struct {
	ImageBase64 string `json:"image_base64"`
	ImageType   string `json:"image_type"`
}

type PolicyAnalysisResponse struct {
	StructuredData string `json:"structured_data"`
	Analysis       string `json:"analysis"`
}
