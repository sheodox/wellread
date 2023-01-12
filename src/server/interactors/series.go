package interactors

import (
	"github.com/sheodox/wellread/repositories"
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

func (s *SeriesInteractor) Add(userId int, name string) (repositories.SeriesEntity, error) {
	if name == "" {
		return repositories.SeriesEntity{}, ErrInvalidName
	}

	return s.repo.Add(userId, name)
}

func (s *SeriesInteractor) Delete(userId, id int) error {
	//todo validate
	s.repo.Delete(userId, id)
	return nil
}

func (s *SeriesInteractor) Update(userId, id int, name, notes string) error {
	if name == "" {
		return ErrInvalidName
	}

	s.repo.Update(userId, id, name, notes)
	return nil
}

func (s *SeriesInteractor) Get(userId, id int) (repositories.SeriesEntity, error) {
	return s.repo.Get(userId, id)
}
