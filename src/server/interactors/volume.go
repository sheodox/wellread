package interactors

import (
	"database/sql"
	"errors"

	"github.com/sheodox/wellread/query"
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

func (v *VolumeInteractor) List(userId int, seriesId sql.NullInt32, name, status sql.NullString, pageNumber int) ([]query.ListVolumesRow, error) {
	if err := v.validateStatus(status.String); status.Valid && err != nil {
		return nil, err
	}
	return v.repo.List(userId, seriesId, name, status, pageNumber)
}

func (v *VolumeInteractor) Get(userId, volumeId int) (query.GetVolumeRow, error) {
	return v.repo.Get(userId, volumeId)
}

func (v *VolumeInteractor) Add(userId, seriesId int, name string) (query.Volume, error) {
	if name == "" {
		return query.Volume{}, ErrInvalidName
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
