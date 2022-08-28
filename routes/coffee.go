package routes

import (
	"goTest/controller"
	"goTest/middleware"
	"goTest/models"
	"goTest/repository"
	"goTest/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CoffeeAPIRoutes(apiRoute *gin.RouterGroup, db *mongo.Database) {

	var (
		coffeeRepository models.CoffeeRepository = repository.NewCoffeeRepository(db)
		coffeeService    models.CoffeeService    = service.NewCoffeeService(coffeeRepository)
		coffeeController models.CoffeeController = controller.NewCoffeeController(coffeeService)
	)

	apiRoute.GET("/test", middleware.JwtAuth(), controller.CoffeeTest)

	apiRoute.POST("/coffee", middleware.JwtAuth(), coffeeController.AddCoffee)

}
