package main

import (
	"AI-Insurance-Agent/config"
	"AI-Insurance-Agent/internal/client"
	"AI-Insurance-Agent/internal/handler"
	"AI-Insurance-Agent/internal/middleware"
	"AI-Insurance-Agent/internal/repository"
	"AI-Insurance-Agent/internal/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/app.yaml")
	if err != nil {
		log.Fatal("配置文件加载失败:", err)
	}

	// 初始化数据库
	if err := config.InitDB(cfg.Database.DSN); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 设置 JWT Secret
	middleware.JWTSecret = []byte(cfg.JWT.Secret)

	// 初始化客户端
	glmClient := &client.GLMClient{
		APIKey: cfg.GLM.APIKey,
		URL:    cfg.GLM.URL,
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

	log.Printf("服务启动在 :%d\n", cfg.Server.Port)
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
