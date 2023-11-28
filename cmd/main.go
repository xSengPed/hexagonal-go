package main

import (
	"context"
	"fmt"
	"hexa-go/configs"
	routes "hexa-go/pkg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	cfg := configs.LoadEnv()
	connectingString := fmt.Sprintf("mongodb://%s:%s@%s/?replicaSet=atlas-oyot2x-shard-0&ssl=true&authSource=admin", cfg.MongoConfig.UserName, cfg.MongoConfig.Password, cfg.MongoConfig.ClusterName)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectingString))
	if err != nil {
		panic(err)
	}
	app := routes.NewRoute(mongoClient).InitializeRouter()
	app.Listen(":3000")
}
