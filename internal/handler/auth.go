package handler

import (
	"AI-Insurance-Agent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": "参数错误"})
		return
	}

	user, err := h.authService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
		"data": gin.H{
			"user_id":  user.ID,
			"username": user.Username,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": "参数错误"})
		return
	}

	token, user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1003, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data": gin.H{
			"token":    token,
			"user_id":  user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
