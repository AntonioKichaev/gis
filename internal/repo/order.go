package repo

import (
	"context"
	"sync"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/repo/model"
)

type Order struct {
	mu      sync.Mutex
	storage []model.Order
}

func NewOrder() *Order {
	return &Order{
		mu: sync.Mutex{},
	}
}

func (o *Order) CreateOrder(ctx context.Context, input *dto.CreateOrderRepoInput) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.storage = append(o.storage, *model.NewOrderFromOrder(input.Entity))

	return nil
}
