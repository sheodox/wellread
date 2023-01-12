package repositories

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/wellread/db"
)

type SeriesEntity struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Notes       string    `db:"notes"`
	CreatedAt   time.Time `db:"created_at"`
	UserId      int       `db:"user_id"`
	VolumeCount int       `db:"volume_count"`
}

type SeriesRepository struct {
	db *sqlx.DB
}

func getUserSeriesQuery(where string) string {
	return "select series.*, count(volumes.id) as volume_count from series left join volumes on volumes.series_id = series.id where series.user_id=$1 " + where + " group by series.id order by name asc"
}

func NewSeriesRepository() *SeriesRepository {
	return &SeriesRepository{db.Connection}
}

func (s *SeriesRepository) List(userId int) ([]SeriesEntity, error) {
	series := []SeriesEntity{}

	err := s.db.Select(&series, getUserSeriesQuery(""), userId)

	fmt.Printf("got this many series %v", len(series))

	return series, err
}

func (s *SeriesRepository) Get(userId, id int) (SeriesEntity, error) {
	series := SeriesEntity{}

	err := s.db.Get(&series, getUserSeriesQuery("and series.id=$2"), userId, id)

	return series, err
}

func (s *SeriesRepository) Add(userId int, name string) (SeriesEntity, error) {
	var id int
	err := s.db.Get(&id, "insert into series (name, created_at, user_id) values ($1, $2, $3) returning id", name, time.Now(), userId)

	if err != nil {
		return SeriesEntity{}, err
	}
	return s.Get(userId, id)
}

func (s *SeriesRepository) Delete(userId, id int) error {
	//todo validate
	_, err := s.db.Exec("delete from series where id=$1 and user_id=$2", id, userId)
	return err
}

func (s *SeriesRepository) Update(userId, id int, name, notes string) error {
	//todo validate
	_, err := s.db.Exec("update series set name=$1, notes=$2 where id=$3 and user_id=$4", name, notes, id, userId)
	return err
}
