package db

import (
	"context"

	"github.com/0xlvl3/workout_api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// How we pass info
type WorkoutStore interface {
	InsertWorkout(context.Context, *types.Workout) (*types.Workout, error)
	GetWorkouts(context.Context, bson.M) ([]*types.Workout, error)
}

// Mongo client and collection init for Workouts
type MongoWorkoutStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// Where we init our store and params
func NewMongoWorkoutStore(client *mongo.Client) *MongoWorkoutStore {
	return &MongoWorkoutStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("workouts"),
	}
}

func (s *MongoWorkoutStore) GetWorkouts(ctx context.Context, filter bson.M) ([]*types.Workout, error) {

	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var workouts []*types.Workout
	if err := resp.All(ctx, &workouts); err != nil {
		return nil, err
	}
	return workouts, nil
}

func (s *MongoWorkoutStore) InsertWorkout(ctx context.Context, workout *types.Workout) (*types.Workout, error) {
	resp, err := s.coll.InsertOne(ctx, workout)
	if err != nil {
		return nil, err
	}

	workout.ID = resp.InsertedID.(primitive.ObjectID)
	return workout, nil
}
