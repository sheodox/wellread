package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/wellread/db"
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

const GET_USER_VOLUME = "select volumes.*, series.name as series_name from volumes inner join series on volumes.series_id = series.id where volumes.user_id=$1 "

type VolumeRepository struct {
	db *sqlx.DB
}

func NewVolumeRepository() *VolumeRepository {
	return &VolumeRepository{db.Connection}
}

func (v *VolumeRepository) Get(userId, volumeId int) (VolumeEntity, error) {
	volume := VolumeEntity{}

	err := v.db.Get(&volume, GET_USER_VOLUME+"and volumes.id=$2", userId, volumeId)

	return volume, err
}

func (v *VolumeRepository) List(userId int) ([]VolumeEntity, error) {
	volumes := []VolumeEntity{}

	err := v.db.Select(&volumes, GET_USER_VOLUME+"order by name asc", userId)

	return volumes, err
}

func (v *VolumeRepository) ListBySeries(userId, seriesId int) ([]VolumeEntity, error) {
	volumes := []VolumeEntity{}

	err := v.db.Select(&volumes, GET_USER_VOLUME+"and series_id=$2 order by name asc", userId, seriesId)

	return volumes, err
}

func (v *VolumeRepository) ListByStatus(userId int, status string) ([]VolumeEntity, error) {
	volumes := []VolumeEntity{}

	err := v.db.Select(&volumes, GET_USER_VOLUME+"and status=$2 order by name asc", userId, status)

	return volumes, err
}

func (v *VolumeRepository) Add(userId, seriesId int, name string) (VolumeEntity, error) {
	volume := VolumeEntity{}
	err := v.db.Get(&volume, "insert into volumes (series_id, name, created_at, user_id) values ($1, $2, $3, $4) returning *", seriesId, name, time.Now(), userId)
	return volume, err
}

func (v *VolumeRepository) Delete(userId, volumeId int) error {
	_, err := v.db.Exec("delete from volumes where id=$1 and user_id=$2", volumeId, userId)

	return err
}

type VolumeEntityUpdateArgs struct {
	Name        string `db:"name"`
	Notes       string `db:"notes"`
	CurrentPage int    `db:"current_page"`
	Status      string `db:"status"`
}

func (v *VolumeRepository) Update(userId, volumeId int, update *VolumeEntityUpdateArgs) error {
	_, err := v.db.Exec("update volumes set notes=$1, current_page=$2, name=$3, status=$4 where id=$5 and user_id=$6", update.Notes, update.CurrentPage, update.Name, update.Status, volumeId, userId)
	return err
}
