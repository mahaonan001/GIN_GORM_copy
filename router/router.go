package router

import (
	"GIN_GORM/controller"
	"GIN_GORM/mail"
	"GIN_GORM/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.GET("/info", middleware.AuthMiddleware(), controller.Info)
		auth.POST("/get_code", mail.Code_email)
	}
	return r
}
