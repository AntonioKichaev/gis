package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/entity"
)

var (
	ErrValidateNewOrder = errors.New("invalid order")
	ErrInvalidDateRange = errors.New("invalid date range")
	ErrUnavailableDays  = errors.New("hotel room is not available for selected dates")
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, input *dto.CreateOrderRepoInput) error
}

type BookingAdapter interface {
	CreateBooking(ctx context.Context, input *dto.CreateBookingInput) error
}

type Order struct {
	r OrderRepo
}

func NewOrder(r OrderRepo) *Order {
	return &Order{
		r: r,
	}
}

func (o *Order) CreateOrder(ctx context.Context, input *dto.CreateOrderInput) error {
	if err := validateNewOrder(input); err != nil {
		return fmt.Errorf("invalid order: %w %v", ErrValidateNewOrder, err)
	}

	e := &entity.Order{
		HotelID:   input.HotelID,
		RoomID:    input.RoomID,
		UserEmail: input.UserEmail,
		From:      input.From,
		To:        input.To,
	}

	if err := o.r.CreateOrder(ctx, &dto.CreateOrderRepoInput{
		Entity: e,
	}); err != nil {
		return fmt.Errorf("o.r.CreateOrder(...): %w", err)
	}

	return nil
}
