package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
	"github.com/sheodox/wellread/repositories"
)

type SeriesController struct {
	interactor *interactors.SeriesInteractor
}

func NewSeriesController() *SeriesController {
	return &SeriesController{&interactors.Series}
}

type seriesRequest struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type seriesResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"createdAt"`
	VolumeCount int       `json:"volumeCount"`
}

func seriesEntityToResponse(entity repositories.SeriesEntity) seriesResponse {
	return seriesResponse{
		Id:          entity.Id,
		Name:        entity.Name,
		Notes:       entity.Notes,
		CreatedAt:   entity.CreatedAt,
		VolumeCount: entity.VolumeCount,
	}
}

func (s *SeriesController) Add(c echo.Context) error {
	body := new(seriesRequest)

	if err := c.Bind(body); err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	entity, err := s.interactor.Add(userId, body.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, seriesEntityToResponse(entity))
}

func (s *SeriesController) Delete(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))

	if err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	s.interactor.Delete(userId, seriesId)

	return c.NoContent(http.StatusOK)
}

func (s *SeriesController) Update(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))

	if err != nil {
		return err
	}

	body := new(seriesRequest)

	if err := c.Bind(body); err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	s.interactor.Update(userId, seriesId, body.Name, body.Notes)

	return c.NoContent(http.StatusOK)
}

func (s *SeriesController) List(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	seriesEntities, err := s.interactor.List(userId)

	if err != nil {
		return err
	}

	series := make([]seriesResponse, len(seriesEntities))

	for i, entity := range seriesEntities {
		series[i] = seriesEntityToResponse(entity)
	}

	return c.JSON(http.StatusOK, series)
}

func (s *SeriesController) Get(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))

	if err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	entity, err := s.interactor.Get(userId, seriesId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, seriesEntityToResponse(entity))
}
