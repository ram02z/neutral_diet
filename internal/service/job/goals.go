package job

import (
	"context"
	"time"

	"github.com/ram02z/neutral_diet/internal/gen/db"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func (j *JobWrapper) markCompletedGoalsJob() {
	queries := db.New(j.p)
	ctx, cancel := context.WithTimeout(
		*j.c,
		// TODO: figure out a reasonable time for the job
		1*time.Minute,
	)
	defer cancel()

	logger := zerolog.Ctx(*j.c)

	userIDs, err := queries.GetUserIDs(ctx)
	if err != nil {
		logger.Err(err).Msg("Could not get user ids")
		return
	}

	completed := 0
	failed := 0
	total := 0
	for _, id := range userIDs {
		dailyAverage, err := queries.GetUserDailyAverageCarbonFootprint(ctx, id)
		if err != nil {
			dailyAverage = decimal.NewFromFloat(0)
		}
		if dailyAverage.IsZero() {
			continue
		}
		goals, err := queries.GetActiveCarbonFootprintGoals(ctx, id)
		if err != nil {
			logger.Err(err).Int32("UserID", id).Msg("Could not get active carbon footprint goals")
			return
		}
		for _, g := range goals {
			if dailyAverage.LessThanOrEqual(g.TargetCarbonFootprint) {
				err := queries.UpdateCarbonFootprintGoal(ctx, db.UpdateCarbonFootprintGoalParams{
					UserID:    id,
					ID:        g.ID,
					Completed: true,
				})
				if err != nil {
					logger.
						Err(err).
						Int32("UserID", id).
						Int32("GoalID", g.ID).
						Msg("Could not update carbon footprint goal")

					failed += 1
				} else {
					completed += 1
				}
			}
			total += 1
		}
	}

	logger.
		Info().
		Int("Completed", completed).
		Int("Failed", failed).
		Int("Total", total).
		Msg("MarkCompletedGoalsJob finished")
}
