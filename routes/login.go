package routes

import (
	"goTest/controller"
	"goTest/repository"
	"goTest/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginRoutes(siteRoute *gin.RouterGroup, db *mongo.Database) {

	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService)
	)

	siteRoute.GET("/login", controller.LoginIndex)

	siteRoute.POST("/login", userController.Login)

}
