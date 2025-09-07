package rooms

import (
	"context"

	"github.com/kalogs-c/gochat/internal/domain"
	sqlc "github.com/kalogs-c/gochat/sql/sqlc_generated"
)

type repository struct {
	db *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) domain.RoomRepository {
	return &repository{q}
}

func (r *repository) CreateRoom(ctx context.Context, room domain.Room) (domain.Room, error) {
	insertedRoom, err := r.db.CreateRoom(ctx, room.Topic)
	if err != nil {
		return domain.Room{}, err
	}

	return domain.Room{
		ID:    insertedRoom.ID,
		Topic: insertedRoom.Topic,
	}, nil
}

func (r *repository) GetRoomByID(ctx context.Context, id int64) (domain.Room, error) {
	room, err := r.db.GetRoom(ctx, id)
	if err != nil {
		return domain.Room{}, err
	}

	return domain.Room{
		ID:    room.ID,
		Topic: room.Topic,
	}, nil
}
