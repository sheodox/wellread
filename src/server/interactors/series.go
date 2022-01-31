package interactors

import (
	"errors"

	"github.com/sheodox/wellread/repositories"
)

var (
	ErrInvalidName = errors.New("Invalid name")
)

type SeriesInteractor struct {
	repo *repositories.SeriesRepository
}

func NewSeriesInteractor() *SeriesInteractor {
	return &SeriesInteractor{repositories.Series}
}

func (s *SeriesInteractor) List(userId int) ([]repositories.SeriesEntity, error) {
	return repositories.Series.List(userId)
}

func (s *SeriesInteractor) Add(userId int, name string) error {
	//todo validate
	s.repo.Add(userId, name)
	return nil
}

func (s *SeriesInteractor) Delete(userId, id int) error {
	//todo validate
	s.repo.Delete(userId, id)
	return nil
}

func (s *SeriesInteractor) Update(userId, id int, name string) error {
	//todo validate
	s.repo.Update(userId, id, name)
	return nil
}
