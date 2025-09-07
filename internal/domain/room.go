package domain

import "context"

type Room struct {
	ID    int64
	Topic string
}

type RoomService interface {
	CreateRoom(ctx context.Context, topic string) (Room, error)
	GetRoomByID(ctx context.Context, id int64) (Room, error)
}

type RoomRepository interface {
	CreateRoom(ctx context.Context, room Room) (Room, error)
	GetRoomByID(ctx context.Context, id int64) (Room, error)
}
