package repository

import (
	"context"
	"goTest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUser(map[string]string) (models.User, error)
}

type database struct {
	connect *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {

	collection := db.Collection("user")

	return &database{
		connect: collection,
	}
}

func (db *database) FindUser(inputData map[string]string) (models.User, error) {

	var result bson.M
	var err error

	err = db.connect.FindOne(context.TODO(), bson.D{}).Decode(&result)

	user := models.User{
		Name:     result["name"].(string),
		Password: result["password"].(string),
		Email:    result["email"].(string),
	}

	return user, err
}
