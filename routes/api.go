package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hailong-bot/phoenix-gin/app/controller/app"
	"github.com/hailong-bot/phoenix-gin/app/controller/common"
	"github.com/hailong-bot/phoenix-gin/app/middleware"
	"github.com/hailong-bot/phoenix-gin/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
}
