package volume

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Volume struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	SeriesId    string    `db:"series_id" json:"seriesId"`
	Notes       string    `json:"notes"`
	CurrentPage int       `db:"current_page" json:"currentPage"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

type VolumeRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *VolumeRepository {
	return &VolumeRepository{db}
}

func (s *VolumeRepository) List(seriesId int) ([]Volume, error) {
	volumes := []Volume{}

	err := s.db.Select(&volumes, "select * from volumes where series_id=$1 order by created_at asc", seriesId)

	return volumes, err
}

func (s *VolumeRepository) Add(seriesId int, name string) {
	s.db.MustExec("insert into volumes (name, series_id, created_at, current_page) values ($1, $2, $3, $4)", name, seriesId, time.Now(), 0)
}

func (s *VolumeRepository) Delete(volumeId int) {
	s.db.MustExec("delete from volumes where id=$1", volumeId)
}

type VolumeUpdate struct {
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	CurrentPage int    `db:"current_page" json:"currentPage"`
}

func (s *VolumeRepository) Update(volumeId int, update *VolumeUpdate) (Volume, error) {
	s.db.MustExec("update volumes set notes=$1, current_page=$2, name=$3 where id=$4", update.Notes, update.CurrentPage, update.Name, volumeId)

	volume := Volume{}

	err := s.db.Get(&volume, "select * from volumes where id=$1", volumeId)

	return volume, err
}
