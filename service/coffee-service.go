package service

import (
	"goTest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type coffeeService struct {
	repository models.CoffeeRepository
}

func NewCoffeeService(coffeeRepository models.CoffeeRepository) models.CoffeeService {
	return &coffeeService{
		repository: coffeeRepository,
	}
}

func (cs *coffeeService) AddCoffee(ctx *gin.Context) {

	price, _ := strconv.Atoi(ctx.PostForm("price"))
	point, _ := strconv.Atoi(ctx.PostForm("point"))

	coffee := &models.Coffee{
		Owner: ctx.PostForm("owner"),
		Type:  ctx.PostForm("type"),
		Price: price,
		Size:  ctx.PostForm("size"),
		Store: models.CoffeeStore{
			Store: ctx.PostForm("store"),
			Point: point,
		},
	}
	cs.repository.AddCoffee(*coffee)
}
