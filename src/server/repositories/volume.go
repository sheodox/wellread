package repositories

import (
	"database/sql"
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

func sqlNullIfBlankString(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid:  str != "",
	}
}
func sqlNullIfBlankInt(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: i != 0,
	}
}

func (v *VolumeRepository) List(userId int, seriesId sql.NullInt32, name, status sql.NullString, pageNumber int) ([]query.ListVolumesRow, error) {
	if name.Valid {
		name.String = "%" + name.String + "%"
	}

	return v.queries.ListVolumes(v.ctx, query.ListVolumesParams{
		UserID:     int32(userId),
		Status:     status,
		Name:       name,
		SeriesID:   seriesId,
		PageOffset: int32(pageNumber * PageSize),
		PageSize:   int32(PageSize),
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
