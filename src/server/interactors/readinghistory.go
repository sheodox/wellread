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

func (r *ReadingHistoryInteractor) Add(userId, volumeId, currentPage int) error {
	return r.repo.Add(userId, volumeId, currentPage)
}

func (r *ReadingHistoryInteractor) List(userId, volumeId int) ([]repositories.ReadingHistoryEntity, error) {
	history, err := r.repo.List(userId, volumeId)

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *ReadingHistoryInteractor) Delete(userId, historyId int) error {
	return r.repo.Delete(userId, historyId)
}
