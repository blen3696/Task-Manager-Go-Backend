package routers

import (
	"github.com/blen/task_manager_api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouters(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.RegisterUser)
		authRouter.POST("/login", controllers.LoginUser)
	}
}
