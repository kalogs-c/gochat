package rooms

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kalogs-c/gochat/internal/domain"
	"github.com/kalogs-c/gochat/pkg/httpjson"
	"github.com/kalogs-c/gochat/pkg/validator"
)

type HTTPAdapter struct {
	service domain.RoomService
	logger  *slog.Logger
}

func NewHTTPAdapter(s domain.RoomService, logger *slog.Logger) *HTTPAdapter {
	return &HTTPAdapter{s, logger}
}

func (h *HTTPAdapter) error(w http.ResponseWriter, r *http.Request, code int, title string, err error) {
	httpjson.NotifyError(
		r.Context(),
		w,
		r,
		h.logger,
		code,
		title,
		err.Error(),
		err,
	)
}

func (h *HTTPAdapter) CreateRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payload, err := httpjson.DecodeValid[*CreateRoomPayload](r)
	if err != nil {
		switch e := err.(type) {
		case validator.ValidationError:
			httpjson.EncodeValidationErrors(w, r, e.Problems)
		default:
			h.error(w, r, http.StatusBadRequest, "invalid payload", err)
		}
		return
	}

	room, err := h.service.CreateRoom(ctx, payload.Topic)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, "failed to create room", err)
		return
	}

	if err := httpjson.Encode(w, r, http.StatusCreated, room); err != nil {
		h.error(w, r, http.StatusInternalServerError, "failed to encode room", err)
	}
}

func (h *HTTPAdapter) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idString := r.PathValue("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		h.error(w, r, http.StatusUnprocessableEntity, "failed to parse id", err)
		return
	}

	room, err := h.service.GetRoomByID(ctx, id)
	if err != nil {
		h.error(w, r, http.StatusNotFound, "failed to retrieve room", err)
		return
	}

	if err := httpjson.Encode(w, r, http.StatusOK, room); err != nil {
		h.error(w, r, http.StatusInternalServerError, "failed to encode room", err)
	}
}

func (h *HTTPAdapter) ListRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.service.ListRooms(r.Context())
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, "failed to list rooms", err)
		return
	}

	if err := httpjson.Encode(w, r, http.StatusOK, rooms); err != nil {
		h.error(w, r, http.StatusInternalServerError, "failed to encode rooms", err)
	}
}
