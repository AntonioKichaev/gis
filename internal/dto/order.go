package dto

import (
	"time"

	"github.com/AntonioKichaev/gis/internal/entity"
)

// todo might be usecase/repo input/output

type CreateOrderInput struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

type CreateOrderRepoInput struct {
	Entity *entity.Order
}
