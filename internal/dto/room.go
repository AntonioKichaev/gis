package dto

import (
	"time"

	"github.com/AntonioKichaev/gis/internal/entity"
)

type CreateBookingRoomInput struct {
	HotelID string
	RoomID  string
	From    time.Time
	To      time.Time
}

type (
	GetAvailabilityRoomInput struct {
		HotelID string
		RoomID  string
	}

	GetAvailabilityRoomOutput struct {
		Entities []entity.RoomAvailability
	}
)

type UpdateAvailabilityRoomInput struct {
	HotelID        string
	RoomID         string
	Availabilities []entity.RoomAvailability
}
