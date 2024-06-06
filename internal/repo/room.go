package repo

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/repo/model"
)

type hotelIDKey string
type roomIDKey string

var (
	ErrHotelNotFound = errors.New("hotel not found")
	ErrRoomNotFound  = errors.New("room not found")
)

type Room struct {
	mu      sync.Mutex
	storage map[hotelIDKey]map[roomIDKey][]model.RoomAvailability
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func NewRoom() *Room {
	storage := make(map[hotelIDKey]map[roomIDKey][]model.RoomAvailability)
	const (
		redisonHotelID = "reddison"
		luxRoomID      = "lux"
	)

	storage[redisonHotelID] = make(map[roomIDKey][]model.RoomAvailability)
	storage[redisonHotelID][luxRoomID] = []model.RoomAvailability{
		{HotelID: redisonHotelID, RoomID: luxRoomID, Date: date(2024, 1, 1), Quota: 1},
		{HotelID: redisonHotelID, RoomID: luxRoomID, Date: date(2024, 1, 2), Quota: 1},
		{HotelID: redisonHotelID, RoomID: luxRoomID, Date: date(2024, 1, 3), Quota: 1},
		{HotelID: redisonHotelID, RoomID: luxRoomID, Date: date(2024, 1, 4), Quota: 1},
		{HotelID: redisonHotelID, RoomID: luxRoomID, Date: date(2024, 1, 5), Quota: 0},
	}

	return &Room{
		storage: storage,
	}
}

func (r *Room) GetAvailability(ctx context.Context, input *dto.GetAvailabilityRoomInput) (*dto.GetAvailabilityRoomOutput, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	hotel, ok := r.storage[hotelIDKey(input.HotelID)]
	if !ok {
		return nil, fmt.Errorf("%w: hotelID: %s", ErrHotelNotFound, input.HotelID)
	}

	room, ok := hotel[roomIDKey(input.RoomID)]
	if !ok {
		return nil, fmt.Errorf("%w: roomID: %s", ErrRoomNotFound, input.RoomID)
	}

	return &dto.GetAvailabilityRoomOutput{
		Entities: model.RoomAvailabilityToEntity(room),
	}, nil
}

func (r *Room) UpdateAvailability(ctx context.Context, input *dto.UpdateAvailabilityRoomInput) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	hotel, ok := r.storage[hotelIDKey(input.HotelID)]
	if !ok {
		return fmt.Errorf("%w: hotelID: %s", ErrHotelNotFound, input.HotelID)
	}

	// todo think about make before update check
	hotel[roomIDKey(input.RoomID)] = model.RoomAvailabilityFromEntities(input.Availabilities)

	return nil
}
