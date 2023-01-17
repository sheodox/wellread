package repositories

import (
	"time"

	"github.com/sheodox/wellread/db"
	"github.com/sheodox/wellread/query"
	"golang.org/x/net/context"
)

type VolumeEntity struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	SeriesId    int       `db:"series_id"`
	UserId      int       `db:"user_id"`
	Notes       string    `db:"notes"`
	CurrentPage int       `db:"current_page"`
	CreatedAt   time.Time `db:"created_at"`
	Status      string    `db:"status"`
	SeriesName  string    `db:"series_name"`
}

type VolumeRepository struct {
	queries *query.Queries
	ctx     context.Context
}

func NewVolumeRepository() *VolumeRepository {
	return &VolumeRepository{db.Queries, context.Background()}
}

func (v *VolumeRepository) Get(userId, volumeId int) (query.GetVolumeRow, error) {
	volume, err := v.queries.GetVolume(v.ctx, query.GetVolumeParams{
		UserID:   int32(userId),
		VolumeID: int32(volumeId),
	})

	return volume, err
}

func (v *VolumeRepository) List(userId int) ([]query.ListVolumesRow, error) {
	return v.queries.ListVolumes(v.ctx, int32(userId))
}

func (v *VolumeRepository) ListBySeries(userId, seriesId int) ([]query.ListVolumesBySeriesRow, error) {
	return v.queries.ListVolumesBySeries(v.ctx, query.ListVolumesBySeriesParams{
		UserID:   int32(userId),
		SeriesID: int32(seriesId),
	})
}

func (v *VolumeRepository) ListByStatus(userId int, status string) ([]query.ListVolumesByStatusRow, error) {
	return v.queries.ListVolumesByStatus(v.ctx, query.ListVolumesByStatusParams{
		UserID: int32(userId),
		Status: status,
	})
}

func (v *VolumeRepository) Add(userId, seriesId int, name string) (query.Volume, error) {
	return v.queries.AddVolume(v.ctx, query.AddVolumeParams{
		UserID:   int32(userId),
		SeriesID: int32(seriesId),
		Name:     name,
	})
}

func (v *VolumeRepository) Delete(userId, volumeId int) error {
	return v.queries.DeleteVolume(v.ctx, query.DeleteVolumeParams{
		UserID:   int32(userId),
		VolumeID: int32(volumeId),
	})
}

type VolumeEntityUpdateArgs struct {
	Name        string
	Notes       string
	CurrentPage int
	Status      string
}

func (v *VolumeRepository) Update(userId, volumeId int, update *VolumeEntityUpdateArgs) error {
	return v.queries.UpdateVolume(v.ctx, query.UpdateVolumeParams{
		UserID:      int32(userId),
		VolumeID:    int32(volumeId),
		Name:        update.Name,
		Notes:       update.Notes,
		CurrentPage: int32(update.CurrentPage),
		Status:      update.Status,
	})
}
