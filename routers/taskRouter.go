package routers

import (
	"github.com/blen/task_manager_api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTaskRouters() *gin.Engine {
	router := gin.Default()

	taskRouter := router.Group("/tasks")
	{
		taskRouter.GET("", controllers.GetTasks)
		taskRouter.GET("/:id", controllers.GetTaskById)
		taskRouter.POST("", controllers.CreateTask)
		taskRouter.PUT("/:id", controllers.UpdateTask)
		taskRouter.DELETE("/:id", controllers.DeleteTask)
	}

	return router
}
