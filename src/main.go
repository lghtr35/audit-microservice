package main

import (
	"audit-backend/container"
	_ "audit-backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title			Audit Micro-Service
// @version         1.0
// @description     This is an audit server for events that are sourced from different services

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")

	}
	g := gin.Default()
	// Initialize Classes
	container.Initialize(g.Group("/api/v1"))
	// Initialize Swagger
	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run()
}
