package job

import (
	"context"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	"github.com/rs/zerolog"
)

type ctxKey struct{}

func (j *JobWrapper) userGoalNotificationsJob() {
	queries := db.New(j.p)

	ctx, cancel := context.WithTimeout(
		*j.c,
		30*time.Second,
	)
	defer cancel()

	logger := zerolog.Ctx(*j.c)

	response, err := queries.GetUserFCMTokens(context.Background())
	if err != nil {
		logger.Err(err).Msg("Could not get user FCM tokens")
		cancel()
	}

	if len(response) == 0 {
		logger.Info().Msg("No user FCM tokens")
	}

	fcmTokens := make([]string, len(response))
	for i, j := range response {
		fcmTokens[i] = j.String
	}

	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title:    "Test",
			Body:     "test body",
		},
		Tokens: fcmTokens,
	}

	br, err := j.m.SendMulticast(ctx, message)
	if err != nil {
		logger.Err(err).Msg("Could not send messages")
	}

	logger.
		Info().
		Int("Success Count", br.SuccessCount).
		Int("Failure Count", br.FailureCount).
		Msg("UserGoalNotificationsJob finished")
}
