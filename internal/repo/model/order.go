package model

import (
	"time"

	"github.com/AntonioKichaev/gis/internal/entity"
)

type Order struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

func NewOrderFromOrder(order *entity.Order) *Order {
	return &Order{
		HotelID:   order.HotelID,
		RoomID:    order.RoomID,
		UserEmail: order.UserEmail,
		From:      order.From,
		To:        order.To,
	}
}
