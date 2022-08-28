package controller

import (
	"goTest/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		service: userService,
	}
}

func LoginIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": "TestSuccess",
	})
}

func (uc *userController) Login(ctx *gin.Context) {
	inputData := make(map[string]string)
	inputData["email"] = ctx.PostForm("email")
	inputData["password"] = ctx.PostForm("password")

	if user, err := uc.service.FindUser(inputData); err == nil {
		//add JWT Token
		token := service.GenerateToken(&user)
		ctx.JSON(200, gin.H{
			"success": token,
		})
	} else {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}

}
