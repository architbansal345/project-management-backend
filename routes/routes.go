package routes

import (
	"project-management-backend/controllers"
	"project-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	public := router.Group("/auth")
	{
		public.POST("/signup", controllers.SignUp)
		public.POST("/login", controllers.Login)
	}
	private := router.Group("/api/v1")
	private.Use(middleware.AuthMiddleware())
	{
		private.POST("/apply-leave", controllers.ApplyLeave)
		private.GET("/leave-balance", controllers.RemainingLeave)
		private.GET("/view-leaveApplications", controllers.ViewLeaveApplication)
	}

}
