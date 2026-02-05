package main

import (
	"AI-Insurance-Agent/internal/client"
	"AI-Insurance-Agent/internal/handler"
	"AI-Insurance-Agent/internal/service"
	"log"
	"net/http"
)

func main() {
	glmClient := &client.GLMClient{
		APIKey: "0ffcfbc67b30483a93a6fa041936f76b.OS41RBK74yt6mskO",
		URL:    "https://open.bigmodel.cn/api/paas/v4/chat/completions",
	}

	policyService, err := service.NewPolicyService(glmClient)
	if err != nil {
		log.Fatal(err)
	}

	policyHandler := handler.NewPolicyHandler(policyService)

	http.HandleFunc("/api/analyze", policyHandler.AnalyzePolicy)

	log.Println("服务启动在 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
