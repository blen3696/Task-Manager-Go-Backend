package main

import (
	"github.com/blen/task_manager_api/routers"
)

func main() {
	router := routers.SetupTaskRouters()
	router.Static("/docs", "./docs")
	router.Run("localhost:8080")
}
