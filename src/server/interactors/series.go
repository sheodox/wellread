package interactors

import (
	"github.com/sheodox/wellread/query"
	"github.com/sheodox/wellread/repositories"
)

type SeriesInteractor struct {
	repo *repositories.SeriesRepository
}

func NewSeriesInteractor() *SeriesInteractor {
	return &SeriesInteractor{repositories.Series}
}

func (s *SeriesInteractor) List(userId int) ([]query.ListSeriesRow, error) {
	return repositories.Series.List(userId)
}

func (s *SeriesInteractor) Add(userId int, name string) (query.GetSeriesRow, error) {
	if name == "" {
		return query.GetSeriesRow{}, ErrInvalidName
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

func (s *SeriesInteractor) Get(userId, id int) (query.GetSeriesRow, error) {
	return s.repo.Get(userId, id)
}
