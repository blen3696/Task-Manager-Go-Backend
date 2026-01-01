package routers

import (
	"github.com/blen/task_manager_api/controllers"
	"github.com/blen/task_manager_api/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRouters(router *gin.Engine) {
	taskRouter := router.Group("/tasks")
	taskRouter.Use(middleware.AuthMiddleware())
	{
		// any authenticated user can get the tasks
		taskRouter.GET("", controllers.GetTasks)
		taskRouter.GET("/:id", controllers.GetTaskById)

		// only admin can perform this operations
		taskRouter.POST("", middleware.AdminOnly(), controllers.CreateTask)
		taskRouter.PUT("/:id", middleware.AdminOnly(), controllers.UpdateTask)
		taskRouter.DELETE("/:id", middleware.AdminOnly(), controllers.DeleteTask)
	}
}
