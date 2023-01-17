package repositories

import (
	"context"
	"time"

	"github.com/sheodox/wellread/db"
	"github.com/sheodox/wellread/query"
)

type SeriesRepository struct {
	queries *query.Queries
	ctx     context.Context
}

func getUserSeriesQuery(where string) string {
	return "select series.*, count(volumes.id) as volume_count from series left join volumes on volumes.series_id = series.id where series.user_id=$1 " + where + " group by series.id order by name asc"
}

func NewSeriesRepository() *SeriesRepository {
	return &SeriesRepository{db.Queries, context.Background()}
}

func (s *SeriesRepository) List(userId int) ([]query.ListSeriesRow, error) {
	return s.queries.ListSeries(s.ctx, int32(userId))
}

func (s *SeriesRepository) Get(userId, id int) (query.GetSeriesRow, error) {
	return s.queries.GetSeries(s.ctx, query.GetSeriesParams{
		UserID:   int32(userId),
		SeriesID: int32(id),
	})
}

func (s *SeriesRepository) Add(userId int, name string) (query.GetSeriesRow, error) {
	id, err := s.queries.AddSeries(s.ctx, query.AddSeriesParams{
		Name:      name,
		CreatedAt: time.Now(),
		UserID:    int32(userId),
	})

	if err != nil {
		return query.GetSeriesRow{}, err
	}
	return s.Get(userId, int(id))
}

func (s *SeriesRepository) Delete(userId, id int) error {
	return s.queries.DeleteSeries(s.ctx, query.DeleteSeriesParams{
		SeriesID: int32(id),
		UserID:   int32(userId),
	})
}

func (s *SeriesRepository) Update(userId, id int, name, notes string) error {
	return s.queries.UpdateSeries(s.ctx, query.UpdateSeriesParams{
		Notes:    notes,
		Name:     name,
		SeriesID: int32(id),
		UserID:   int32(userId),
	})
}
