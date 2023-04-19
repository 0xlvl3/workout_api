package main

import (
	"context"
	"flag"
	"log"

	"github.com/0xlvl3/workout_api/api"
	"github.com/0xlvl3/workout_api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "Listen address of API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// handler init
	var (
		userHandle    = api.NewUserHandler(db.NewMongoUserStore(client))
		workoutHandle = api.NewWorkoutHandler(db.NewMongoWorkoutStore(client))
	)

	// fiber init
	var (
		app   = fiber.New()
		apiv1 = app.Group("/api/v1/")
	)

	// user handlers
	apiv1.Get("/user", userHandle.HandleGetUsers)
	apiv1.Get("/user/:id", userHandle.HandleGetUser)

	// workout handlers
	apiv1.Get("/workout", workoutHandle.HandleGetWorkouts)

	app.Listen(*listenAddr)

}
