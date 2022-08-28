package controller

import (
	"goTest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type coffeeController struct {
	service models.CoffeeService
}

func NewCoffeeController(coffeeService models.CoffeeService) models.CoffeeController {
	return &coffeeController{
		service: coffeeService,
	}
}

func CoffeeTest(ctx *gin.Context) {

	var coffees = []models.Coffee{
		{Owner: "ian", Type: "black", Price: 35, Size: "middle", Store: models.CoffeeStore{Store: "familyMart", Point: 10}},
		{Owner: "ian", Type: "latte", Price: 45, Size: "middle", Store: models.CoffeeStore{Store: "familyMart", Point: 20}},
	}

	ctx.IndentedJSON(http.StatusOK, coffees)

}

func (cc *coffeeController) AddCoffee(ctx *gin.Context) {
	var coffee models.Coffee

	if err := ctx.Bind(&coffee); err == nil {
		cc.service.AddCoffee(ctx)
		ctx.JSON(200, gin.H{"message": "Success"})
	} else {
		ctx.JSON(200, gin.H{"Error": err.Error()})
	}

}
