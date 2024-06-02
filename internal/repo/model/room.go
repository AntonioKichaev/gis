package model

import (
	"time"

	"github.com/AntonioKichaev/gis/internal/entity"
)

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}

func RoomAvailabilityToEntity(ra []RoomAvailability) []entity.RoomAvailability {
	result := make([]entity.RoomAvailability, len(ra))

	for i := range ra {
		result[i] = *toEntity(&ra[i])
	}

	return result
}

func toEntity(r *RoomAvailability) *entity.RoomAvailability {
	return &entity.RoomAvailability{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
		Date:    r.Date,
		Quota:   r.Quota,
	}
}

func RoomAvailabilityFromEntities(ra []entity.RoomAvailability) []RoomAvailability {
	result := make([]RoomAvailability, len(ra))

	for i := range ra {
		result[i] = *roomAvailabilityFromEntity(&ra[i])
	}

	return result
}

func roomAvailabilityFromEntity(r *entity.RoomAvailability) *RoomAvailability {
	return &RoomAvailability{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
		Date:    r.Date,
		Quota:   r.Quota,
	}
}
