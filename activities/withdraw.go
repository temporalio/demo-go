package activities

import (
	"context"

	"go.temporal.io/temporal/activity"
	"go.uber.org/zap"
)

func Withdraw(ctx context.Context, accountId, referenceId string, amount int) error {
	logger := activity.GetLogger(ctx)
	logger.Info("withdrawal requested",
		zap.String("AccountId", accountId),
		zap.String("ReferenceId", referenceId),
		zap.Int("Amount", amount))

	return nil
}
