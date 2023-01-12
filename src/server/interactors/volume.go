package interactors

import (
	"errors"

	"github.com/sheodox/wellread/repositories"
)

var (
	validStatuses    = []string{"planning", "reading", "completed", "dropped"}
	ErrInvalidStatus = errors.New("invalid status")
	ErrInvalidPages  = errors.New("invalid current page")
)

type VolumeInteractor struct {
	repo              *repositories.VolumeRepository
	historyInteractor *ReadingHistoryInteractor
}

func NewVolumeInteractor() *VolumeInteractor {
	return &VolumeInteractor{repositories.Volume, &ReadingHistory}
}

func (v *VolumeInteractor) validateStatus(status string) error {
	statusValid := false
	for _, aValidStatus := range validStatuses {
		if aValidStatus == status {
			statusValid = true
		}
	}

	if statusValid {
		return nil
	}
	return ErrInvalidStatus
}

func (v *VolumeInteractor) ListBySeries(userId, seriesId int) ([]repositories.VolumeEntity, error) {
	return v.repo.ListBySeries(userId, seriesId)
}

func (v *VolumeInteractor) List(userId int) ([]repositories.VolumeEntity, error) {
	return v.repo.List(userId)
}

func (v *VolumeInteractor) Get(userId, volumeId int) (repositories.VolumeEntity, error) {
	return v.repo.Get(userId, volumeId)
}

func (v *VolumeInteractor) ListByStatus(userId int, status string) ([]repositories.VolumeEntity, error) {
	if err := v.validateStatus(status); err != nil {
		return nil, err
	}

	return v.repo.ListByStatus(userId, status)
}

func (v *VolumeInteractor) Add(userId, seriesId int, name string) (repositories.VolumeEntity, error) {
	if name == "" {
		return repositories.VolumeEntity{}, ErrInvalidName
	}

	return v.repo.Add(userId, seriesId, name)
}

func (v *VolumeInteractor) Delete(userId, volumeId int) {
	v.repo.Delete(userId, volumeId)
}

type VolumeUpdateArgs struct {
	Name        string
	Notes       string
	CurrentPage int
	Status      string
	PagesRead   int
}

func (v *VolumeInteractor) Update(userId, volumeId int, update *VolumeUpdateArgs) error {
	if update.CurrentPage < 0 {
		return ErrInvalidPages
	}

	if update.Name == "" {
		return ErrInvalidName
	}

	if err := v.validateStatus(update.Status); err != nil {
		return err
	}

	if update.PagesRead != 0 {
		err := v.historyInteractor.Add(userId, volumeId, update.CurrentPage, update.PagesRead)

		if err != nil {
			return err
		}

	}

	err := v.repo.Update(userId, volumeId, &repositories.VolumeEntityUpdateArgs{
		CurrentPage: update.CurrentPage,
		Notes:       update.Notes,
		Name:        update.Name,
		Status:      update.Status,
	})

	return err
}
