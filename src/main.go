package main

import (
	"audit-backend/container"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")

	}
	g := gin.Default()
	// Initialize Classes
	container.Initialize(g.Group("/api"))
	// Initialize DB Connection

	g.Run()
}
