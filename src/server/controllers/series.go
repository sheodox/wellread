package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
)

type SeriesController struct {
	interactor *interactors.SeriesInteractor
}

func NewSeriesController() *SeriesController {
	return &SeriesController{&interactors.Series}
}

type seriesRequest struct {
	Name string `json:"name"`
}

type seriesResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
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

	s.interactor.Add(userId, body.Name)

	return s.List(c)
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

	return s.List(c)
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

	s.interactor.Update(userId, seriesId, body.Name)

	return s.List(c)
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
		series[i] = seriesResponse{
			Id:        entity.Id,
			Name:      entity.Name,
			Notes:     entity.Notes,
			CreatedAt: entity.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, series)
}
