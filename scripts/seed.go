package main

import (
	"context"
	"fmt"
	"log"

	"github.com/0xlvl3/workout_api/db"
	"github.com/0xlvl3/workout_api/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	workoutStore db.WorkoutStore
	ctx          = context.Background()
)

func seedWorkout(name string, numOfEx int32) {

	workout := &types.Workout{
		Name:              name,
		NumberOfExercises: numOfEx,
		//	ExercisesList:     []primitive.ObjectID{},
	}

	insertedWorkout, err := workoutStore.InsertWorkout(ctx, workout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedWorkout)
}

func main() {
	seedWorkout("Push", 12)
	seedWorkout("Pull", 10)
	seedWorkout("Legs", 20)
}

func init() {
	var err error

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	workoutStore = db.NewMongoWorkoutStore(client)
}
