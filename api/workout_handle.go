package api

import (
	"github.com/0xlvl3/workout_api/db"
	"github.com/gofiber/fiber/v2"
)

type WorkoutHandler struct {
	workoutStore db.WorkoutStore
}

func NewWorkoutHandler(ws db.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: ws,
	}
}

func (w *WorkoutHandler) HandleGetWorkouts(c *fiber.Ctx) error {
	workouts, err := w.workoutStore.GetWorkouts(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(workouts)
}
