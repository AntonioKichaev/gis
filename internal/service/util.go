package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/AntonioKichaev/gis/internal/dto"
)

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)

	rightLimit := toDay(to)
	for d := toDay(from); !d.After(rightLimit); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func validateNewOrder(newOrder *dto.CreateOrderInput) error {
	var errs []error

	if newOrder.HotelID == "" {
		errs = append(errs, fmt.Errorf("hotel_id is required"))
	}

	if newOrder.RoomID == "" {
		errs = append(errs, fmt.Errorf("room_id is required"))
	}

	if newOrder.UserEmail == "" {
		errs = append(errs, fmt.Errorf("email is required"))
	}

	if newOrder.From.IsZero() {
		errs = append(errs, fmt.Errorf("from is required"))
	}

	if newOrder.To.IsZero() {
		errs = append(errs, fmt.Errorf("to is required"))
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w", errors.Join(errs...))
	}

	return nil
}
