package rooms

import (
	"context"

	"github.com/kalogs-c/gochat/pkg/validator"
)

type RoomResponse struct {
	ID    int64  `json:"id"`
	Topic string `json:"topic"`
}

type CreateRoomPayload struct {
	Topic string `json:"topic"`
}

func (crp *CreateRoomPayload) Valid(ctx context.Context) validator.Problems {
	problems := make(validator.Problems)

	if len(crp.Topic) < 3 {
		problems["topic"] = append(problems["topic"], "should be at least 3 chars long")
	}

	return problems
}
