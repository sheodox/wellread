package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/wellread/db"
)

type VolumeEntity struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	SeriesId    string    `db:"series_id"`
	Notes       string    `db:"notes"`
	CurrentPage int       `db:"current_page"`
	CreatedAt   time.Time `db:"created_at"`
}

type VolumeRepository struct {
	db *sqlx.DB
}

func (v *VolumeRepository) FindOne(volumeId int) (VolumeEntity, error) {
	volume := VolumeEntity{}

	err := v.db.Get(&volume, "select * from volumes where id=$1", volumeId)

	return volume, err
}

func NewVolumeRepository() *VolumeRepository {
	return &VolumeRepository{db.Connection}
}

func (v *VolumeRepository) List(seriesId int) ([]VolumeEntity, error) {
	volumes := []VolumeEntity{}

	err := v.db.Select(&volumes, "select * from volumes where series_id=$1 order by created_at asc", seriesId)

	return volumes, err
}

func (v *VolumeRepository) Add(seriesId int, name string) error {
	_, err := v.db.Exec("insert into volumes (series_id, name, created_at) values ($1, $2, $3)", seriesId, name, time.Now())
	return err
}

func (v *VolumeRepository) Delete(volumeId int) error {
	_, err := v.db.Exec("delete from volumes where id=$1", volumeId)

	return err
}

type VolumeEntityUpdateArgs struct {
	Name        string `db:"name"`
	Notes       string `db:"notes"`
	CurrentPage int    `db:"current_page"`
}

func (v *VolumeRepository) Update(volumeId int, update *VolumeEntityUpdateArgs) error {
	_, err := v.db.Exec("update volumes set notes=$1, current_page=$2, name=$3 where id=$4", update.Notes, update.CurrentPage, update.Name, volumeId)
	return err
}
