package main

import (
	"audit-backend/container"
	_ "audit-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Audit Micro-Service
// @version         1.0
// @description     This is an audit server for events that are sourced from different services

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	// TODO concurrency design could be better
	// TODO CQRS might be a better solution for performance
	g := gin.Default()
	// Initialize Application
	container.Initialize(g.Group("/api/v1"))
	// Initialize Swagger
	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Start Server
	container.StartServer(g)
}
