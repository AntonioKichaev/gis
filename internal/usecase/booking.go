package usecase

import (
	"context"
	"fmt"

	"github.com/AntonioKichaev/gis/internal/dto"
)

//go:generate mockery --name RoomServiceAdapter --case=underscore --with-expecter
type RoomServiceAdapter interface {
	CreateBooking(ctx context.Context, input *dto.CreateBookingRoomInput) error
}

//go:generate mockery --name OrderServiceAdapter --case=underscore --with-expecter
type OrderServiceAdapter interface {
	CreateOrder(ctx context.Context, input *dto.CreateOrderInput) error
}

type Booking struct {
	r RoomServiceAdapter
	o OrderServiceAdapter
}

func NewBooking(r RoomServiceAdapter, o OrderServiceAdapter) *Booking {
	return &Booking{
		r: r,
		o: o,
	}
}

func (b *Booking) CreateBooking(ctx context.Context, input *dto.CreateBookingInput) error {
	// todo must be in transaction by tr manager

	if err := b.r.CreateBooking(ctx, &dto.CreateBookingRoomInput{
		HotelID: input.HotelID,
		RoomID:  input.RoomID,
		From:    input.From,
		To:      input.To,
	}); err != nil {
		return fmt.Errorf("o.b.CreateBooking(...): %w", err)
	}

	if err := b.o.CreateOrder(ctx, &dto.CreateOrderInput{
		HotelID:   input.HotelID,
		RoomID:    input.RoomID,
		UserEmail: input.UserEmail,
		From:      input.From,
		To:        input.To,
	}); err != nil {
		return fmt.Errorf("o.o.CreateOrder(...): %w", err)
	}

	return nil
}
