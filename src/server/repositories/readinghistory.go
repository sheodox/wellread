package repositories

import (
	"context"
	"time"

	"github.com/sheodox/wellread/db"
	"github.com/sheodox/wellread/query"
)

type ReadingHistoryEntity struct {
	Id          int       `db:"id"`
	VolumeId    int       `db:"volume_id"`
	CurrentPage int       `db:"current_page"`
	PagesRead   int       `db:"pages_read"`
	CreatedAt   time.Time `db:"created_at"`
	UserId      int       `db:"user_id"`
}

type ReadingHistoryRepository struct {
	queries *query.Queries
	ctx     context.Context
}

func NewReadingHistoryRepository() *ReadingHistoryRepository {
	return &ReadingHistoryRepository{db.Queries, context.Background()}
}

func (r *ReadingHistoryRepository) List(userId, volumeId int) ([]query.ReadingHistory, error) {
	return r.queries.ListReadingHistory(r.ctx, query.ListReadingHistoryParams{
		UserID:   int32(userId),
		VolumeID: int32(volumeId),
	})
}

func (r *ReadingHistoryRepository) Add(userId, volumeId, currentPage, pagesRead int) error {
	return r.queries.AddReadingHistory(r.ctx, query.AddReadingHistoryParams{
		VolumeID:    int32(volumeId),
		UserID:      int32(userId),
		CreatedAt:   time.Now(),
		CurrentPage: int32(currentPage),
		PagesRead:   int32(pagesRead),
	})
}

func (r *ReadingHistoryRepository) Delete(userId, id int) error {
	return r.queries.DeleteReadingHistory(r.ctx, query.DeleteReadingHistoryParams{
		ReadingHistoryID: int32(id),
		UserID:           int32(userId),
	})
}
