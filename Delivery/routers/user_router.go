package routers

import (
	"github.com/blen/task_manager_api/Delivery/controllers"
	"github.com/blen/task_manager_api/Infrastructure"

	"github.com/gin-gonic/gin"
)

func UserRoutes(
	r *gin.Engine,
	controller *controllers.UserController,
	jwt *infrastructure.JWTService,
) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	admin := r.Group("/admin")
	admin.Use(infrastructure.AuthMiddleware(jwt))
	{
		admin.PUT("/promote/:id", controller.Promote)
	}
}
