package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
)

type createVolumeRequest struct {
	Name string `json:"name"`
}

type VolumeController struct {
	interactor *interactors.VolumeInteractor
}

func NewVolumeController() *VolumeController {
	return &VolumeController{&interactors.Volume}
}

func (v *VolumeController) Add(c echo.Context) error {
	body := new(createVolumeRequest)

	if err := c.Bind(body); err != nil {
		return err
	}

	seriesId, err := strconv.Atoi(c.Param("seriesId"))
	if err != nil {
		return err
	}

	v.interactor.Add(seriesId, body.Name)

	return v.List(c)
}

type volumeUpdateRequest struct {
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	CurrentPage int    `json:"currentPage"`
}

func (v *VolumeController) Update(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	body := new(volumeUpdateRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	err = v.interactor.Update(volumeId, &interactors.VolumeUpdateArgs{
		Name:        body.Name,
		Notes:       body.Notes,
		CurrentPage: body.CurrentPage,
	})

	if err != nil {
		return err
	}

	return v.List(c)
}

func (v *VolumeController) Delete(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	v.interactor.Delete(volumeId)

	if err != nil {
		return err
	}

	return v.List(c)
}

type volumeResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Notes       string    `json:"notes"`
	CurrentPage int       `json:"currentPage"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (v *VolumeController) List(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))
	if err != nil {
		return err
	}

	volumeEntities, err := v.interactor.List(seriesId)

	if err != nil {
		return err
	}

	volumes := make([]volumeResponse, len(volumeEntities))

	for i, entity := range volumeEntities {
		volumes[i] = volumeResponse{
			Id:          entity.Id,
			Name:        entity.Name,
			CurrentPage: entity.CurrentPage,
			Notes:       entity.Notes,
		}
	}

	return c.JSON(http.StatusOK, volumes)
}
