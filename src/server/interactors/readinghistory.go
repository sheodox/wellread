package interactors

import (
	"github.com/sheodox/wellread/repositories"
)

type ReadingHistoryInteractor struct {
	repo repositories.ReadingHistoryRepository
}

func NewReadingHistoryInteractor() *ReadingHistoryInteractor {
	return &ReadingHistoryInteractor{*repositories.ReadingHistory}
}

func (r *ReadingHistoryInteractor) Add(volumeId, currentPage int) error {
	return r.repo.Add(volumeId, currentPage)
}

func (r *ReadingHistoryInteractor) List(volumeId int) ([]repositories.ReadingHistoryEntity, error) {
	history, err := r.repo.List(volumeId)

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *ReadingHistoryInteractor) Delete(historyId int) error {
	return r.repo.Delete(historyId)
}
