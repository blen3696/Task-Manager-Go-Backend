package routers

import (
	"github.com/blen/task_manager_api/controllers"
	"github.com/blen/task_manager_api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouters(router *gin.Engine) {
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.PUT("/promote/:id", controllers.PromoteUser)
	}
}
