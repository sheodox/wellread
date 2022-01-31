package interactors

import (
	"github.com/sheodox/wellread/repositories"
)

type VolumeInteractor struct {
	repo              *repositories.VolumeRepository
	historyInteractor *ReadingHistoryInteractor
}

func NewVolumeInteractor() *VolumeInteractor {
	return &VolumeInteractor{repositories.Volume, &ReadingHistory}
}

func (v *VolumeInteractor) List(userId, seriesId int) ([]repositories.VolumeEntity, error) {
	return v.repo.List(userId, seriesId)
}

func (v *VolumeInteractor) Add(userId, seriesId int, name string) {
	v.repo.Add(userId, seriesId, name)
}

func (v *VolumeInteractor) Delete(userId, volumeId int) {
	v.repo.Delete(userId, volumeId)
}

type VolumeUpdateArgs struct {
	Name        string
	Notes       string
	CurrentPage int
}

func (v *VolumeInteractor) Update(userId, volumeId int, update *VolumeUpdateArgs) error {
	existingVolume, err := v.repo.FindOne(userId, volumeId)

	if err != nil {
		return err
	}

	if update.CurrentPage != existingVolume.CurrentPage {
		err = v.historyInteractor.Add(userId, volumeId, update.CurrentPage)

		if err != nil {
			return err
		}

	}

	err = v.repo.Update(userId, volumeId, &repositories.VolumeEntityUpdateArgs{
		CurrentPage: update.CurrentPage,
		Notes:       update.Notes,
		Name:        update.Name,
	})

	return err
}
