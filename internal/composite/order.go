package composite

import (
	"github.com/AntonioKichaev/gis/internal/handler"
	"github.com/AntonioKichaev/gis/internal/repo"
	"github.com/AntonioKichaev/gis/internal/service"
	"github.com/AntonioKichaev/gis/internal/usecase"
)

type Order struct {
	Handler *handler.Order
}

func NewOrder(l handler.Logger) *Order {
	bookingRepo := repo.NewRoom()
	roomSvc := service.NewRoom(bookingRepo)

	orderRepo := repo.NewOrder()
	orderSvc := service.NewOrder(orderRepo)

	uc := usecase.NewBooking(roomSvc, orderSvc)

	return &Order{
		Handler: handler.NewOrder(l, uc),
	}
}
