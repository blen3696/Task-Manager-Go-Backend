package controllers

import (
	// "net/http"

	"github.com/blen/task_manager_api/Usecases"
	"github.com/blen/task_manager_api/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewUserController(u *usecases.UserUsecase) *UserController {
	return &UserController{u}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user domain.User
	ctx.ShouldBindJSON(&user)
	token, err := c.usecase.Register(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"token": token})
}

func (c *UserController) Login(ctx *gin.Context) {
	var input struct {
		Email    string
		Password string
	}
	ctx.ShouldBindJSON(&input)
	token, err := c.usecase.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

func (c *UserController) Promote(ctx *gin.Context) {
	id := ctx.Param("id")
	c.usecase.Promote(id)
	ctx.JSON(200, gin.H{"message": "user promoted"})
}
