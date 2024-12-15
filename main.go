package main

import (
	"fmt"
	"project-management-backend/config"
	"project-management-backend/middleware"
	"project-management-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r)
	r.Run(":8080")
	fmt.Println("hello")
}
