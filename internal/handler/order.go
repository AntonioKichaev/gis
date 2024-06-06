package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/service"
)

type Logger interface {
	LogErrorf(format string, v ...any)
	LogInfo(format string, v ...any)
}

type OrderAdapter interface {
	CreateBooking(ctx context.Context, input *dto.CreateBookingInput) error
}

type Order struct {
	l Logger
	u OrderAdapter
}

func NewOrder(l Logger, u OrderAdapter) *Order {
	return &Order{
		l: l,
		u: u,
	}
}

type OrderCreateRequest struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (h *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		h.l.LogErrorf("Invalid request: %s", err)

		return
	}

	if err := h.u.CreateBooking(r.Context(), &dto.CreateBookingInput{
		HotelID:   newOrder.HotelID,
		RoomID:    newOrder.RoomID,
		UserEmail: newOrder.UserEmail,
		From:      newOrder.From,
		To:        newOrder.To,
	}); err != nil {
		setErrorStatusCode(w, err)

		h.l.LogErrorf("h.u.CreateBooking(...): %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newOrder); err != nil {
		h.l.LogErrorf("Failed to encode order: %v", err)
	}

	h.l.LogInfo("Order successfully created: %v", newOrder)
}

func setErrorStatusCode(w http.ResponseWriter, err error) {
	if errors.Is(err, service.ErrValidateNewOrder) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrInvalidDateRange) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrUnavailableDays) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
