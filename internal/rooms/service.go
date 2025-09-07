package rooms

import (
	"context"

	"github.com/kalogs-c/gochat/internal/domain"
)

type service struct {
	repository domain.RoomRepository
}

func NewService(roomRepository domain.RoomRepository) domain.RoomService {
	return &service{roomRepository}
}

func (s *service) CreateRoom(ctx context.Context, topic string) (domain.Room, error) {
	room := domain.Room{Topic: topic}
	return s.repository.CreateRoom(ctx, room)
}

func (s *service) GetRoomByID(ctx context.Context, id int64) (domain.Room, error) {
	return s.repository.GetRoomByID(ctx, id)
}
