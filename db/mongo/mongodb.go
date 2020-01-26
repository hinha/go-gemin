package mongo

import (
	"fmt"
	"context"
	"go-gemin/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Name of the database.
	DBName = "glottery"

)
var URI = "mongodb://" + config.Option.MongoUser + ":"+ config.Option.MongoPass +"@"+ config.Option.MongoURL +"/" + config.Option.MongoDB
var DB *mongo.Database
func init() {
	ctx := context.Background()

	clientOption := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println(err)
		return
	}

	DB = client.Database(config.Option.MongoDB)

}