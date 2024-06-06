package dto

import "time"

type CreateBookingInput struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}
