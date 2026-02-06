package handler

import (
	"AI-Insurance-Agent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PolicyHandler struct {
	service *service.PolicyService
}

func NewPolicyHandler(service *service.PolicyService) *PolicyHandler {
	return &PolicyHandler{service: service}
}

func (h *PolicyHandler) AnalyzePolicy(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req struct {
		ImageBase64 string `json:"image_base64" binding:"required"`
		ImageType   string `json:"image_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": "参数错误"})
		return
	}

	record, err := h.service.AnalyzePolicy(userID, req.ImageBase64, req.ImageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2001, "message": "分析失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "分析成功",
		"data":    record,
	})
}

func (h *PolicyHandler) GetRecord(c *gin.Context) {
	userID := c.GetInt64("user_id")
	recordID, _ := strconv.ParseInt(c.Param("record_id"), 10, 64)

	record, err := h.service.GetRecord(recordID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 1002, "message": "记录不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": record})
}

func (h *PolicyHandler) ListRecords(c *gin.Context) {
	userID := c.GetInt64("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	records, total, err := h.service.ListRecords(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
			"records":   records,
		},
	})
}

func (h *PolicyHandler) DeleteRecord(c *gin.Context) {
	userID := c.GetInt64("user_id")
	recordID, _ := strconv.ParseInt(c.Param("record_id"), 10, 64)

	err := h.service.DeleteRecord(recordID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}
