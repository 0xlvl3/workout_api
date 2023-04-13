package main

import (
	"context"
	"flag"
	"log"

	"github.com/0xlvl3/workout_api/api"
	"github.com/0xlvl3/workout_api/db"
	"github.com/0xlvl3/workout_api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	listenAddr := flag.String("listenAddr", ":8080", "Listen address of API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		Username: "Corey",
	}
	coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	apiv1 := app.Group("/api/v1/")

	// handler init
	userHandle := api.NewUserHandler(db.NewMongoUserStore(client))

	apiv1.Get("/user", userHandle.HandleGetUsers)
	apiv1.Get("/user/:id", userHandle.HandleGetUser)

	app.Listen(*listenAddr)

}
