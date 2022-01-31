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
	UserId    int       `db:"user_id"`
}

type SeriesRepository struct {
	db *sqlx.DB
}

func NewSeriesRepository() *SeriesRepository {
	return &SeriesRepository{db.Connection}
}

func (s *SeriesRepository) List(userId int) ([]SeriesEntity, error) {
	series := []SeriesEntity{}

	err := s.db.Select(&series, "select * from series where user_id=$1 order by name asc", userId)

	return series, err
}

func (s *SeriesRepository) Add(userId int, name string) error {
	//todo validate
	_, err := s.db.Exec("insert into series (name, created_at, user_id) values ($1, $2, $3)", name, time.Now(), userId)

	return err
}

func (s *SeriesRepository) Delete(userId, id int) error {
	//todo validate
	_, err := s.db.Exec("delete from series where id=$1 and user_id=$2", id, userId)
	return err
}

func (s *SeriesRepository) Update(userId, id int, name string) error {
	//todo validate
	_, err := s.db.Exec("update series set name=$1 where id=$2 and user_id=$3", name, id, userId)
	return err
}
