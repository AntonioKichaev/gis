package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AntonioKichaev/gis/internal/dto"
	"github.com/AntonioKichaev/gis/internal/usecase/mocks"
)

func TestBooking_CreateBooking(t *testing.T) {
	r := mocks.NewRoomServiceAdapter(t)
	o := mocks.NewOrderServiceAdapter(t)

	expected := &dto.CreateBookingInput{
		HotelID:   "1",
		RoomID:    "2", // TODO: better use gofakeit lib
		UserEmail: "123@mail.ru",
		From:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		To:        time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
	}

	r.EXPECT().CreateBooking(mock.AnythingOfType("backgroundCtx"), mock.MatchedBy(
		func(input *dto.CreateBookingRoomInput) bool {
			return assert.Equal(t, expected.HotelID, input.HotelID) &&
				assert.Equal(t, expected.RoomID, input.RoomID) &&
				assert.Equal(t, expected.From, input.From) &&
				assert.Equal(t, expected.To, input.To)

		},
	)).Return(nil).Once()

	o.EXPECT().CreateOrder(mock.AnythingOfType("backgroundCtx"), mock.MatchedBy(
		func(input *dto.CreateOrderInput) bool {
			return assert.Equal(t, expected.HotelID, input.HotelID) &&
				assert.Equal(t, expected.RoomID, input.RoomID) &&
				assert.Equal(t, expected.UserEmail, input.UserEmail) &&
				assert.Equal(t, expected.From, input.From) &&
				assert.Equal(t, expected.To, input.To)

		},
	)).Return(nil).Once()

	b := NewBooking(r, o)

	ctx := context.Background()
	err := b.CreateBooking(ctx, expected)

	assert.NoError(t, err, "Error in CreateBooking() function")

}
