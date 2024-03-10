package reports

import (
	"context"
	"github.com/lrmnt/AA6_homework/analytics/ent"
	"github.com/lrmnt/AA6_homework/analytics/ent/billingoperations"
	"github.com/lrmnt/AA6_homework/analytics/ent/task"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	log    *zap.Logger
	client *ent.Client
}

func New(log *zap.Logger, client *ent.Client) *Service {
	return &Service{log: log, client: client}
}

func (s *Service) GetLastDaySum(ctx context.Context) (int64, error) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)

	ops, err := s.client.BillingOperations.Query().
		Where(billingoperations.And(
			billingoperations.TimestampLTE(now.UnixNano()),
			billingoperations.TimestampGT(yesterday.UnixNano()),
		)).
		All(ctx)
	if err != nil {
		return 0, err
	}

	var out int64
	for i := range ops {
		switch ops[i].Type {
		case billingoperations.TypePayout:
			out -= ops[i].Amount
		case billingoperations.TypeDebtMove:
			out += ops[i].Amount
		}
	}

	return out, nil
}

func (s *Service) GetLastDayPopugDebtsCount(ctx context.Context) (int64, error) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)

	ops, err := s.client.BillingOperations.Query().
		Where(billingoperations.And(
			billingoperations.TimestampLTE(now.UnixNano()),
			billingoperations.TimestampGT(yesterday.UnixNano()),
		)).
		All(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	for i := range ops {
		switch ops[i].Type {
		case billingoperations.TypeDebtMove:
			if ops[i].Amount > 0 {
				count++
			}
		}
	}

	return count, nil
}

func (s *Service) GetHighestPriceForPeriod(ctx context.Context, start, end time.Time) (int64, error) {
	tasks, err := s.client.Task.Query().
		Where(task.And(
			task.TimestampLT(end.UnixNano()),
			task.TimestampGTE(start.UnixNano()),
		)).
		All(ctx)
	if err != nil {
		return 0, err
	}

	var maxPrice int64
	for i := range tasks {
		maxPrice = max(tasks[i].Cost, maxPrice)
	}

	return maxPrice, nil
}
