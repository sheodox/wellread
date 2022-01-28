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

func (v *VolumeInteractor) List(seriesId int) ([]repositories.VolumeEntity, error) {
	return v.repo.List(seriesId)
}

func (v *VolumeInteractor) Add(seriesId int, name string) {
	v.repo.Add(seriesId, name)
}

func (v *VolumeInteractor) Delete(volumeId int) {
	v.repo.Delete(volumeId)
}

type VolumeUpdateArgs struct {
	Name        string
	Notes       string
	CurrentPage int
}

func (v *VolumeInteractor) Update(volumeId int, update *VolumeUpdateArgs) error {
	existingVolume, err := v.repo.FindOne(volumeId)

	if err != nil {
		return err
	}

	if update.CurrentPage != existingVolume.CurrentPage {
		err = v.historyInteractor.Add(volumeId, update.CurrentPage)

		if err != nil {
			return err
		}

	}

	err = v.repo.Update(volumeId, &repositories.VolumeEntityUpdateArgs{
		CurrentPage: update.CurrentPage,
		Notes:       update.Notes,
		Name:        update.Name,
	})

	return err
}
