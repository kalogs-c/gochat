package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/kalogs-c/gochat/internal/rooms"
	sqlc "github.com/kalogs-c/gochat/sql/sqlc_generated"
)

func setupRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	queries *sqlc.Queries,
) {
	setupRooms(mux, logger, queries)
}

func setupRooms(
	mux *http.ServeMux,
	logger *slog.Logger,
	queries *sqlc.Queries,
) {
	repo := rooms.NewRepository(queries)
	service := rooms.NewService(repo)
	adapter := rooms.NewHTTPAdapter(service, logger)

	mux.HandleFunc("GET /rooms", adapter.ListRooms)
	mux.HandleFunc("GET /rooms/{id}", adapter.GetRoomByID)
	mux.HandleFunc("POST /rooms", adapter.CreateRoom)
}
