package nosql

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbDefault() *mongo.Database {
	var (
		client *mongo.Client
		err    error
		db     *mongo.Database
	)

	clientOption := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", viper.Get("mongo.host"), viper.Get("mongo.port")))

	if client, err = mongo.Connect(context.TODO(), clientOption); err != nil {
		panic(err)
	}

	db = client.Database(fmt.Sprintf("%s", viper.Get("mongo.database")))
	return db
}
