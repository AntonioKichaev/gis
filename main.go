// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/AntonioKichaev/gis/config"
	"github.com/AntonioKichaev/gis/internal/composite"
	"github.com/AntonioKichaev/gis/pkg/logger"
)

func main() {
	mux := http.NewServeMux()

	cfg := config.NewConfig()
	l := logger.NewLogger()
	order := composite.NewOrder(l)

	mux.HandleFunc("/orders", order.Handler.CreateOrder)

	l.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(cfg.GetAddr(), mux)
	if errors.Is(err, http.ErrServerClosed) {
		l.LogInfo("Server closed")
	} else if err != nil {
		l.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
