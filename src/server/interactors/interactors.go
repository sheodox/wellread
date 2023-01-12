package interactors

import (
	"errors"
)

var (
	ErrInvalidName = errors.New("Invalid name")
)

var (
	Auth           = *NewAuthInteractor()
	Series         = *NewSeriesInteractor()
	Volume         = *NewVolumeInteractor()
	ReadingHistory = *NewReadingHistoryInteractor()
)
