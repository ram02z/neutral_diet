package db

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func (s *Store) SendGoalNotifications(ctx context.Context, m *messaging.Client) {
	jobName := "SendGoalNotifications"
	queries := db.New(s.dbPool)
	logger := zerolog.Ctx(ctx)

	response, err := queries.GetDevices(context.Background())
	if err != nil {
		logger.Err(err).Str("Job", jobName).Msg("Could not get user FCM tokens")
		return
	}

	if len(response) == 0 {
		logger.Error().Str("Job", jobName).Msg("No user FCM tokens")
		return
	}

	messages := make([]*messaging.Message, len(response))
	for i, j := range response {
		notification, err := generateGoalNotification(ctx, queries, j.UserID)
		if err != nil {
			logger.
				Err(err).
				Str("Job", jobName).
				Int32("UserID", j.UserID).
				Msg("Could not generate notification")
			return
		}
		messages[i] = &messaging.Message{
			Notification: notification,
			Token:        j.FcmToken,
		}
	}

	// TODO: check that messages is < 500
	br, err := m.SendAll(ctx, messages)
	if err != nil {
		logger.Err(err).Msg("Could not send messages")
		return
	}

	logger.
		Info().
		Str("Job", jobName).
		Int("Success Count", br.SuccessCount).
		Int("Failure Count", br.FailureCount).
		Msg("Job finished")
}

func (s *Store) MarkCompletedGoals(ctx context.Context) {
	jobName := "MarkCompletedGoals"
	queries := db.New(s.dbPool)
	logger := zerolog.Ctx(ctx)

	userIDs, err := queries.GetUserIDs(ctx)
	if err != nil {
		logger.Err(err).Str("Job", jobName).Msg("Could not get user ids")
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
			logger.
				Err(err).
				Str("Job", jobName).
				Int32("UserID", id).
				Msg("Could not get active carbon footprint goals")
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
						Str("Job", jobName).
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
		Str("Job", jobName).
		Int("Completed", completed).
		Int("Failed", failed).
		Int("Total", total).
		Msg("Job finished")
}

func (s *Store) SendStreakNotifications(ctx context.Context, m *messaging.Client) {
	jobName := "SendStreakNotifications"
	queries := db.New(s.dbPool)
	logger := zerolog.Ctx(ctx)

	response, err := queries.GetDevices(context.Background())
	if err != nil {
		logger.Err(err).Str("Job", jobName).Msg("Could not get user FCM tokens")
		return
	}

	if len(response) == 0 {
		logger.Error().Str("Job", jobName).Msg("No user FCM tokens")
		return
	}

	messages := make([]*messaging.Message, 0, len(response))
	for _, j := range response {
		notification, err := generateStreakNotification(ctx, queries, j.UserID)
		if err != nil {
			logger.
				Err(err).
				Str("Job", jobName).
				Int32("UserID", j.UserID).
				Msg("Could not generate notification")
			return
		}
		if notification != nil {
			messages = append(messages,
				&messaging.Message{
					Notification: notification,
					Token:        j.FcmToken,
				})
		}
	}

	successCount := 0
	failureCount := 0
	if len(messages) > 0 {
		// TODO: check that messages is < 500
		br, err := m.SendAll(ctx, messages)
		if err != nil {
			logger.Err(err).Msg("Could not send messages")
			return
		}
		successCount = br.SuccessCount
		failureCount = br.FailureCount
	}

	logger.
		Info().
		Str("Job", jobName).
		Int("Success Count", successCount).
		Int("Failure Count", failureCount).
		Msg("Job finished")
}

func generateStreakNotification(
	ctx context.Context,
	queries *db.Queries,
	userID int32,
) (*messaging.Notification, error) {
	todayLog, err := queries.GetFoodItemLogByDate(ctx, db.GetFoodItemLogByDateParams{
		UserID: userID,
		LogDate: pgtype.Date{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(todayLog) > 0 {
		return nil, nil
	}

	streak, err := queries.GetFoodItemLogStreak(ctx, userID)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	if !streak.Active {
		return nil, nil
	}

	return &messaging.Notification{
		Title: "Streak reminder",
		Body: fmt.Sprintf(
			"Ensure you log food today to maintain your %d day streak!",
			streak.ConsecutiveDates,
		),
	}, nil
}

func generateGoalNotification(
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

	todayLog, err := queries.GetFoodItemLogByDate(ctx, db.GetFoodItemLogByDateParams{
		UserID: userID,
		LogDate: pgtype.Date{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	notification := messaging.Notification{}

	if userDailyAverage.IsZero() && len(activeGoals) == 0 {
		notification.Title = "New user reminder"
		notification.Body = "Try adding a carbon footprint goal to get started!"
	} else {
		notification.Title = "Carbon footprint goal reminder"
		notification.Body = fmt.Sprintf("You have %d active goals.", len(activeGoals))
		if len(todayLog) == 0 {
			notification.Body += "\nEnsure you log your food before the end of today!"
		} else {
			notification.Body += "\nKeep up the good work!"
		}
	}

	return &notification, nil
}
