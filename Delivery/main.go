package main

import (
	"github.com/blen/task_manager_api/Delivery/controllers"
	"github.com/blen/task_manager_api/Delivery/routers"
	"github.com/blen/task_manager_api/Infrastructure"
	"github.com/blen/task_manager_api/Usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := infrastructure.ConnectMongo()

	userRepo := infrastructure.NewMongoUserRepository(db.Collection("users"))
	passwordSvc := &infrastructure.BcryptPasswordService{}
	jwtSvc := &infrastructure.JWTService{}

	userUsecase := usecases.NewUserUsecase(userRepo, passwordSvc, jwtSvc)
	userController := controllers.NewUserController(userUsecase)

	r := gin.Default()
	routers.UserRoutes(r, userController, jwtSvc)

	r.Run(":8080")
}
