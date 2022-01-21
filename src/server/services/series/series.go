package series

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	seriesRepo "github.com/sheodox/bookmark/repositories/series"
	volumeRepo "github.com/sheodox/bookmark/repositories/volume"
)

type SeriesService struct {
	seriesRepo *seriesRepo.SeriesRepository
	volumeRepo *volumeRepo.VolumeRepository
}

func New(conn *sqlx.DB) *SeriesService {
	return &SeriesService{seriesRepo.New(conn), volumeRepo.New(conn)}
}

type createSeries struct {
	Name string `json:"name"`
}

func (s *SeriesService) Add(c echo.Context) error {
	body := new(createSeries)

	if err := c.Bind(body); err != nil {
		return err
	}

	s.seriesRepo.Add(body.Name)

	return s.List(c)
}

func (s *SeriesService) Delete(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))

	if err != nil {
		return err
	}

	s.seriesRepo.Delete(seriesId)

	return s.List(c)
}

func (s *SeriesService) Update(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))

	if err != nil {
		return err
	}

	body := new(createSeries)

	if err := c.Bind(body); err != nil {
		return err
	}

	s.seriesRepo.Update(seriesId, body.Name)

	return s.List(c)
}

func (s *SeriesService) List(c echo.Context) error {
	series, err := s.seriesRepo.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, series)
}

type createVolume struct {
	Name string `json:"name"`
}

func (s *SeriesService) AddVolume(c echo.Context) error {
	body := new(createVolume)

	if err := c.Bind(body); err != nil {
		return err
	}

	seriesId, err := strconv.Atoi(c.Param("seriesId"))
	if err != nil {
		return err
	}

	s.volumeRepo.Add(seriesId, body.Name)

	return s.ListVolumes(c)
}

func (s *SeriesService) UpdateVolume(c echo.Context) error {
	body := new(volumeRepo.VolumeUpdate)
	if err := c.Bind(body); err != nil {
		return err
	}

	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	_, err = s.volumeRepo.Update(volumeId, body)

	if err != nil {
		return err
	}

	return s.ListVolumes(c)
}

func (s *SeriesService) DeleteVolume(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	s.volumeRepo.Delete(volumeId)

	if err != nil {
		return err
	}

	return s.ListVolumes(c)
}

func (s *SeriesService) ListVolumes(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))
	if err != nil {
		return err
	}

	series, err := s.volumeRepo.List(seriesId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, series)
}
