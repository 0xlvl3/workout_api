package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workout struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"workoutName" json:"workoutName"`
	NumberOfExercises int32              `bson:"numberOfExercises" json:"numberOfExercises"`
	//ExercisesList     []primitive.ObjectID `bson:"exercisesList" json:"exercisesList"`
}

type ExerciseType int

const (
	_            ExerciseType = iota // 0
	ChestType                        // 1
	BackType                         // 2
	ShoulderType                     // 3
	LegType                          // 4
	ArmType                          // 5
)

type Exercise struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	WorkoutID primitive.ObjectID `bson:"workoutID" json:"workoutID"`
	Type      ExerciseType       `bson:"type" json:"type"`
	Weight    float64            `bson:"weight" json:"weight"`
	Rep       int32              `bson:"rep" json:"rep"`
}
