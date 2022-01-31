package repositories

var (
	Series         = NewSeriesRepository()
	Volume         = NewVolumeRepository()
	ReadingHistory = NewReadingHistoryRepository()
	Auth           = NewAuthRepository()
)
