package main

import (
	"AI-Insurance-Agent/config"
	"AI-Insurance-Agent/internal/client"
	"AI-Insurance-Agent/internal/handler"
	"AI-Insurance-Agent/internal/middleware"
	"AI-Insurance-Agent/internal/repository"
	"AI-Insurance-Agent/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/insurance_agent?charset=utf8mb4&parseTime=True&loc=Local"
	if err := config.InitDB(dsn); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化客户端
	glmClient := &client.GLMClient{
		APIKey: "0ffcfbc67b30483a93a6fa041936f76b.OS41RBK74yt6mskO",
		URL:    "https://open.bigmodel.cn/api/paas/v4/chat/completions",
	}

	// 初始化仓储
	userRepo := repository.NewUserRepository(config.DB)
	recordRepo := repository.NewRecordRepository(config.DB)

	// 初始化服务
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	policyService, err := service.NewPolicyService(glmClient, recordRepo)
	if err != nil {
		log.Fatal("PolicyService初始化失败:", err)
	}

	// 初始化处理器
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	policyHandler := handler.NewPolicyHandler(policyService)

	// 设置路由
	r := gin.Default()

	// 公开接口
	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)

	// 需要认证的接口
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		// 用户接口
		auth.GET("/user/profile", userHandler.GetProfile)
		auth.PUT("/user/password", userHandler.ChangePassword)

		// 保单分析接口
		auth.POST("/policy/analyze", policyHandler.AnalyzePolicy)
		auth.GET("/policy/records", policyHandler.ListRecords)
		auth.GET("/policy/records/:record_id", policyHandler.GetRecord)
		auth.DELETE("/policy/records/:record_id", policyHandler.DeleteRecord)

		// 管理员接口
		auth.GET("/admin/users", userHandler.ListUsers)
		auth.PUT("/admin/users/:user_id/status", userHandler.UpdateUserStatus)
	}

	log.Println("服务启动在 :8080")
	r.Run(":8080")
}
