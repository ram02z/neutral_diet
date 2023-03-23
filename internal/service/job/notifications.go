package job

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func generateNotification(
	ctx context.Context,
	queries *db.Queries,
	userID int32,
) (*messaging.Notification, error) {
	userDailyAverage, err := queries.GetUserDailyAverageCarbonFootprint(ctx, userID)
	if err != nil {
		userDailyAverage = decimal.NewFromFloat(0)
	}
	activeGoals, err := queries.GetActiveCarbonFootprintGoals(ctx, userID)
	if err != nil {
		return nil, err
	}

	notification := messaging.Notification{}

	if userDailyAverage.IsZero() && len(activeGoals) == 0 {
		notification.Title = "New user reminder"
		notification.Body = "Try adding a carbon footprint goal to get started!"
	} else {
		notification.Title = "Carbon footprint goal reminder"
		notification.Body = fmt.Sprintf("You have %d active goals", len(activeGoals))
	}

	return &notification, nil
}

func (j *JobWrapper) userGoalNotificationsJob() {
	queries := db.New(j.p)

	ctx, cancel := context.WithTimeout(
		*j.c,
		// TODO: figure out a reasonable time for the job
		1*time.Minute,
	)
	defer cancel()

	logger := zerolog.Ctx(*j.c)

	response, err := queries.GetDevices(context.Background())
	if err != nil {
		logger.Err(err).Msg("Could not get user FCM tokens")
		return
	}

	if len(response) == 0 {
		logger.Info().Msg("No user FCM tokens")
	}

	messages := make([]*messaging.Message, len(response))
	for i, j := range response {
		notification, err := generateNotification(ctx, queries, j.ID)
		if err != nil {
			logger.Err(err).Int32("UserID", j.ID).Msg("Could not generate notification")
			return
		}
		messages[i] = &messaging.Message{
			Notification: notification,
			Token:        j.FcmToken,
		}
	}

	// TODO: check that messages is < 500
	br, err := j.m.SendAll(ctx, messages)
	if err != nil {
		logger.Err(err).Msg("Could not send messages")
		return
	}

	logger.
		Info().
		Int("Success Count", br.SuccessCount).
		Int("Failure Count", br.FailureCount).
		Msg("UserGoalNotificationsJob finished")
}
