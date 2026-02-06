package handler

import (
	"AI-Insurance-Agent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 1002, "message": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": user})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": "参数错误"})
		return
	}

	err := h.userService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1003, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "密码修改成功"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"code": 1005, "message": "权限不足"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total, err := h.userService.ListUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"total": total,
			"users": users,
		},
	})
}

func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"code": 1005, "message": "权限不足"})
		return
	}

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": "参数错误"})
		return
	}

	err := h.userService.UpdateUserStatus(userID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户状态已更新"})
}
