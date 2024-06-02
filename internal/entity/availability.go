package entity

import "time"

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}

func (r *RoomAvailability) IsAvailable() bool {
	return r.Quota > 0
}

func (r *RoomAvailability) DecreaseQuota() {
	r.Quota--
}
