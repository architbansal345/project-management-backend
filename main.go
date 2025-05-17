package main

import (
	"fmt"
	"log"
	"project-management-backend/config"
	"project-management-backend/middleware"
	"project-management-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r)
	r.Run(":8080")
	fmt.Println("hello")
}
