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

func (s *SeriesInteractor) List() ([]repositories.SeriesEntity, error) {
	return repositories.Series.List()
}

func (s *SeriesInteractor) Add(name string) error {
	//todo validate
	s.repo.Add(name)
	return nil
}

func (s *SeriesInteractor) Delete(id int) error {
	//todo validate
	s.repo.Delete(id)
	return nil
}

func (s *SeriesInteractor) Update(id int, name string) error {
	//todo validate
	s.repo.Update(id, name)
	return nil
}
