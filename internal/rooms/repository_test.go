package rooms

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/kalogs-c/gochat/internal/domain"
	"github.com/kalogs-c/gochat/internal/storage/sqlite"
	"github.com/kalogs-c/gochat/sql/migrations"
	sqlc "github.com/kalogs-c/gochat/sql/sqlc_generated"
	"github.com/stretchr/testify/require"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := sqlite.MustConnect(ctx, ":memory:")
	gooseProvider := migrations.MustProvide(db)
	testQueries = sqlc.New(db)

	_, err := gooseProvider.Up(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}

func TestRepository_CreateAndGetRoom(t *testing.T) {
	repo := NewRepository(testQueries)
	ctx := context.Background()

	room, err := repo.CreateRoom(ctx, domain.Room{Topic: "Go Chat"})
	require.NoError(t, err)
	require.NotZero(t, room.ID)

	got, err := repo.GetRoomByID(ctx, room.ID)
	require.NoError(t, err)
	require.Equal(t, "Go Chat", got.Topic)
	require.Equal(t, room.ID, got.ID)
}
