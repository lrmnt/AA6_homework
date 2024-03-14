package schema

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/lib/api/schema/billing_event"
	"github.com/lrmnt/AA6_homework/lib/api/schema/task_event"
	"github.com/lrmnt/AA6_homework/lib/api/schema/task_stream"
	"github.com/lrmnt/AA6_homework/lib/api/schema/user_stream"
	"strings"
)

var (
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	ErrInvalidType      = errors.New("invalid event type")
	ErrNotEnoughData    = errors.New("not enough data")
	ErrWrongData        = errors.New("wrong data")
)

func ValidateTaskEventV1(e *task_event.TaskEventV1) (bool, error) {
	_, err := uuid.Parse(e.EventId)
	if err != nil {
		return false, fmt.Errorf("can not parse event id: %w", err)
	}

	switch e.Event {
	case task_event.Event_EVENT_REASSIGNED, task_event.Event_EVENT_DONE:
	default:
		return false, ErrInvalidType
	}

	_, err = uuid.Parse(e.TaskId)
	if err != nil {
		return false, fmt.Errorf("can not parse task id: %w", err)
	}

	_, err = uuid.Parse(e.AssigneeUserId)
	if err != nil {
		return false, fmt.Errorf("can not parse user id: %w", err)
	}

	if e.Timestamp <= 0 {
		return false, ErrInvalidTimestamp
	}

	return true, nil
}

func ValidateTaskStreamV1(s *task_stream.TaskStreamV1) (bool, error) {
	_, err := uuid.Parse(s.IdempotencyKey)
	if err != nil {
		return false, fmt.Errorf("can not parse idempotency key: %w", err)
	}

	_, err = uuid.Parse(s.PublicId)
	if err != nil {
		return false, fmt.Errorf("can not parse task id: %w", err)
	}

	_, err = uuid.Parse(s.UserId)
	if err != nil {
		return false, fmt.Errorf("can not parse user id: %w", err)
	}

	if s.Timestamp <= 0 {
		return false, ErrInvalidTimestamp
	}

	if s.Title == "" {
		return false, ErrNotEnoughData
	}

	if strings.ContainsAny(s.Title, "[]") {
		return false, ErrWrongData
	}

	if s.Action == task_stream.Action_ACTION_UNKNOWN {
		return false, ErrInvalidType
	}

	return true, nil
}

func ValidateUserStreamV1(s *user_stream.UserStreamV1) (bool, error) {
	_, err := uuid.Parse(s.IdempotencyKey)
	if err != nil {
		return false, fmt.Errorf("can not parse idempotency key: %w", err)
	}

	_, err = uuid.Parse(s.PublicId)
	if err != nil {
		return false, fmt.Errorf("can not parse task id: %w", err)
	}

	_, err = uuid.Parse(s.PublicId)
	if err != nil {
		return false, fmt.Errorf("can not parse user id: %w", err)
	}

	if s.Timestamp <= 0 {
		return false, ErrInvalidTimestamp
	}

	if s.Name == "" || s.Role == "" {
		return false, ErrNotEnoughData
	}

	if s.Action == user_stream.Action_ACTION_UNKNOWN {
		return false, ErrInvalidType
	}

	return true, nil
}

func ValidateBillingEventV1(e *billing_event.BillingEventV1) (bool, error) {
	_, err := uuid.Parse(e.EventId)
	if err != nil {
		return false, fmt.Errorf("can not parse event id: %w", err)
	}

	switch e.Event {
	case billing_event.Event_EVENT_PAYED_TO_USER, billing_event.Event_EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY:
	default:
		return false, ErrInvalidType
	}

	_, err = uuid.Parse(e.UserId)
	if err != nil {
		return false, fmt.Errorf("can not parse user id: %w", err)
	}

	if e.Timestamp <= 0 {
		return false, ErrInvalidTimestamp
	}

	if e.Amount < 0 {
		return false, ErrWrongData
	}

	return true, nil
}
