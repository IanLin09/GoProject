package models

import "github.com/gin-gonic/gin"

type CoffeeStore struct {
	Store string `json:"store"`
	Point int    `json:"point"`
}

type Coffee struct {
	Owner string      `json:"owner"`
	Type  string      `json:"type"`
	Price int         `json:"price"`
	Size  string      `json:"size"`
	Store CoffeeStore `json:"storeInfo" gorm:"embedded"`
}

type CoffeeRepository interface {
	AddCoffee(Coffee)
}

type CoffeeService interface {
	AddCoffee(ctx *gin.Context)
}

type CoffeeController interface {
	AddCoffee(ctx *gin.Context)
}
