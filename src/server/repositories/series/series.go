package series

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Series struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type SeriesRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *SeriesRepository {
	return &SeriesRepository{db}
}

func (s *SeriesRepository) List() ([]Series, error) {
	series := []Series{}

	err := s.db.Select(&series, "select * from series order by name asc")

	return series, err
}

func (s *SeriesRepository) Add(name string) {
	s.db.MustExec("insert into series (name, created_at) values ($1, $2)", name, time.Now())
}

func (s *SeriesRepository) Delete(id int) {
	s.db.MustExec("delete from series where id=$1", id)
}

func (s *SeriesRepository) Update(id int, name string) {
	s.db.MustExec("update series set name=$1 where id=$2", name, id)
}
