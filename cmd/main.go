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
	cfg := configs.NewConfig().InitConfig()
	ctx := context.Background()

	connectingString := fmt.Sprintf("mongodb://%s:%s@%s/?replicaSet=atlas-oyot2x-shard-0&ssl=true&authSource=admin", cfg.MongoConfig.UserName, cfg.MongoConfig.Password, cfg.MongoConfig.ClusterName)

	fmt.Println(connectingString)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectingString))
	if err != nil {
		panic(err)
	}
	app := routes.NewRoute(mongoClient, cfg).InitializeRouter()
	app.Listen(":3000")
}
