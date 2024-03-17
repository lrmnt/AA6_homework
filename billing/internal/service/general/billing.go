package general

import (
	"context"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent"
	"github.com/lrmnt/AA6_homework/billing/ent/billingcycle"
	"time"
)

func GetCurrentBillingCycle(ctx context.Context, tx *ent.Tx) (*ent.BillingCycle, error) {
	curTime := time.Now()

	var fromStartOfDay time.Duration
	fromStartOfDay += time.Hour * time.Duration(curTime.Hour())
	fromStartOfDay += time.Minute * time.Duration(curTime.Minute())
	fromStartOfDay += time.Second * time.Duration(curTime.Second())
	fromStartOfDay += time.Nanosecond * time.Duration(curTime.Nanosecond())

	tsFrom := curTime.Add(-fromStartOfDay).UnixNano()
	tsTo := curTime.Add(-fromStartOfDay).Add(time.Hour * 24).UnixNano()

	bc, err := tx.BillingCycle.Query().
		Where(billingcycle.And(
			billingcycle.TsTo(tsTo),
			billingcycle.TsFrom(tsFrom),
		)).
		Only(ctx)

	if err == nil {
		return bc, nil
	}

	if !ent.IsNotFound(err) {
		return nil, err
	}

	// create new bc
	bc, err = tx.BillingCycle.Create().
		SetUUID(uuid.New()).
		SetTsTo(tsTo).
		SetTsFrom(tsFrom).
		SetState(billingcycle.StateOpen).
		Save(ctx)

	return bc, err
}
