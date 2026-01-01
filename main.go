package main

import (
	"github.com/blen/task_manager_api/db"
	"github.com/blen/task_manager_api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.Default()

	routers.TaskRouters(router)
	routers.AuthRouters(router)
	routers.AdminRouters(router)

	router.Static("/docs", "./docs")
	router.Run("localhost:8080")
}
