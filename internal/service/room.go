package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/entity"
)

type RoomRepo interface {
	GetAvailability(ctx context.Context, input *dto.GetAvailabilityRoomInput) (*dto.GetAvailabilityRoomOutput, error)
	UpdateAvailability(ctx context.Context, input *dto.UpdateAvailabilityRoomInput) error
}

type Room struct {
	r RoomRepo
}

func NewRoom(r RoomRepo) *Room {
	return &Room{
		r: r,
	}
}

func (b *Room) CreateBooking(ctx context.Context, input *dto.CreateBookingRoomInput) error {
	daysToBook := daysBetween(input.From, input.To)
	if len(daysToBook) == 0 {
		return ErrInvalidDateRange
	}

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	output, err := b.r.GetAvailability(ctx, &dto.GetAvailabilityRoomInput{
		HotelID: input.HotelID,
		RoomID:  input.RoomID,
	})
	if err != nil {
		return fmt.Errorf("b.r.GetAvailability(...): %w", err)
	}

	availabilityDateMap := make(map[time.Time]entity.RoomAvailability, len(output.Entities))
	for _, availability := range output.Entities {
		availabilityDateMap[availability.Date] = availability
	}

	for _, dayToBook := range daysToBook {
		availability, ok := availabilityDateMap[dayToBook] // TODO: is it good idea to use dayToBook as key? time.Equal??
		if !ok || !availability.IsAvailable() {
			continue
		}

		availability.DecreaseQuota()
		availabilityDateMap[availability.Date] = availability

		delete(unavailableDays, dayToBook)
	}

	if len(unavailableDays) != 0 {
		return fmt.Errorf("%w: unavailableDays: %v", ErrUnavailableDays, unavailableDays)
	}

	toUpdate := make([]entity.RoomAvailability, 0, len(availabilityDateMap))

	for _, availability := range availabilityDateMap {
		toUpdate = append(toUpdate, availability)
	}

	if err := b.r.UpdateAvailability(ctx, &dto.UpdateAvailabilityRoomInput{
		HotelID:        input.HotelID,
		RoomID:         input.RoomID,
		Availabilities: toUpdate,
	}); err != nil {
		return fmt.Errorf("b.r.UpdateAvailability(...): %w", err)
	}

	return nil
}
