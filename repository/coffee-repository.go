package repository

import (
	"context"
	"goTest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type coffeetable struct {
	connect *mongo.Collection
}

type coffeeRMDB struct {
	db *gorm.DB
}

func NewCoffeeRepository(db *mongo.Database) models.CoffeeRepository {

	collection := db.Collection("coffee")

	return &coffeetable{
		connect: collection,
	}
}

func NewCoffeeRMDBRepository(db *gorm.DB) models.CoffeeRepository {

	return &coffeeRMDB{
		db: db,
	}
}

func (ct *coffeetable) AddCoffee(data models.Coffee) {

	insert := bson.D{
		{"owner", data.Owner},
		{"type", data.Type},
		{"price", data.Price},
		{"size", data.Size},
		{"store", data.Store},
	}

	_, err := ct.connect.InsertOne(context.TODO(), insert)
	if err != nil {
		panic(err)
	}
}

func (cr *coffeeRMDB) AddCoffee(data models.Coffee) {

	result := cr.db.Create(&data)

	if result.Error != nil {
		panic(result.Error)
	}
}
