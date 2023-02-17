package repositories

var (
	Series         = NewSeriesRepository()
	Volume         = NewVolumeRepository()
	ReadingHistory = NewReadingHistoryRepository()
	Auth           = NewAuthRepository()
	PageSize       = 10
)
