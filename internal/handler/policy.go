package handler

import (
	"AI-Insurance-Agent/internal/model"
	"AI-Insurance-Agent/internal/service"
	"encoding/json"
	"net/http"
)

type PolicyHandler struct {
	service *service.PolicyService
}

func NewPolicyHandler(service *service.PolicyService) *PolicyHandler {
	return &PolicyHandler{service: service}
}

func (h *PolicyHandler) AnalyzePolicy(w http.ResponseWriter, r *http.Request) {
	var req model.PolicyAnalysisRequest
	json.NewDecoder(r.Body).Decode(&req)

	result, err := h.service.AnalyzePolicy(req.ImageBase64, req.ImageType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.PolicyAnalysisResponse{
		StructuredData: result,
		Analysis:       result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
