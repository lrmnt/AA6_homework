package reports

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent"
	"github.com/lrmnt/AA6_homework/billing/ent/billing"
	"github.com/lrmnt/AA6_homework/billing/ent/billingcycle"
	"github.com/lrmnt/AA6_homework/billing/ent/operations"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
	"github.com/lrmnt/AA6_homework/billing/internal/service/general"
	"go.uber.org/zap"
)

type Service struct {
	log    *zap.Logger
	client *ent.Client
}

func New(log *zap.Logger, client *ent.Client) *Service {
	return &Service{log: log, client: client}
}

func (s *Service) ListUserBillingLog(ctx context.Context, userID uuid.UUID) ([]*ent.Billing, error) {
	return s.client.Billing.Query().
		Where(billing.HasUserWith(user.UUID(userID))).
		Order(billing.ByTimestamp(sql.OrderDesc())).
		All(ctx)
}

func (s *Service) ListUserOperationsLog(ctx context.Context, userID uuid.UUID) (map[string][]ent.Operations, error) {
	ops, err := s.client.Operations.Query().
		Where(operations.HasUserWith(user.UUID(userID))).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("can not list operations: %w", err)
	}

	out := map[string][]ent.Operations{}
	for i := range ops {
		date := ops[i].Timestamp.Format("2006-01-02")

		sl, _ := out[date]
		sl = append(sl, *ops[i])
		out[date] = sl

	}

	return out, nil
}

func (s *Service) GetUserBalance(ctx context.Context, userID uuid.UUID) (int64, error) {
	u, err := s.client.User.Query().
		Where(user.UUID(userID)).
		Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("can not query user: %w", err)
	}

	return u.Balance, nil
}

type AdminStatisticsDay struct {
	AssignedTasksFeeSum  int64 `json:"assigned_tasks_fee_sum"`
	CompletedTasksAmount int64 `json:"completed_tasks_amount"`
}

type AdminStatisticsReport struct {
	TodayInfo    AdminStatisticsDay            `json:"today_info"`
	PreviousDays map[string]AdminStatisticsDay `json:"previous_days"`
}

func (s *Service) GetAdminStatistics(ctx context.Context) (*AdminStatisticsReport, error) {
	out := AdminStatisticsReport{
		PreviousDays: make(map[string]AdminStatisticsDay),
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	currentBC, err := general.GetCurrentBillingCycle(ctx, tx)
	if err != nil {
		return nil, err
	}

	todayOps, err := tx.Operations.Query().
		Where(operations.HasBillingCycleWith(billingcycle.ID(currentBC.ID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	for i := range todayOps {
		switch todayOps[i].Type {
		case operations.TypeIncome:
			out.TodayInfo.CompletedTasksAmount += todayOps[i].Amount
		case operations.TypeOutcome:
			out.TodayInfo.AssignedTasksFeeSum += todayOps[i].Amount
		}
	}

	// todo other days

	return &out, nil
}
