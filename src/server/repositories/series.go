package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/wellread/db"
)

type SeriesEntity struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Notes     string    `db:"notes"`
	CreatedAt time.Time `db:"created_at"`
}

type SeriesRepository struct {
	db *sqlx.DB
}

func NewSeriesRepository() *SeriesRepository {
	return &SeriesRepository{db.Connection}
}

func (s *SeriesRepository) List() ([]SeriesEntity, error) {
	series := []SeriesEntity{}

	err := s.db.Select(&series, "select * from series order by name asc")

	return series, err
}

func (s *SeriesRepository) Add(name string) error {
	//todo validate
	_, err := s.db.Exec("insert into series (name, created_at) values ($1, $2)", name, time.Now())

	return err
}

func (s *SeriesRepository) Delete(id int) error {
	//todo validate
	_, err := s.db.Exec("delete from series where id=$1", id)
	return err
}

func (s *SeriesRepository) Update(id int, name string) error {
	//todo validate
	_, err := s.db.Exec("update series set name=$1 where id=$2", name, id)
	return err
}
