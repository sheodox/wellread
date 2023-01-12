package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
	"github.com/sheodox/wellread/repositories"
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

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	volumeEntity, err := v.interactor.Add(userId, seriesId, body.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, volumeEntityToResponse(volumeEntity))
}

type volumeUpdateRequest struct {
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	CurrentPage int    `json:"currentPage"`
	PagesRead   int    `json:"pagesRead"`
	Status      string `json:"status"`
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

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	err = v.interactor.Update(userId, volumeId, &interactors.VolumeUpdateArgs{
		Name:        body.Name,
		Notes:       body.Notes,
		CurrentPage: body.CurrentPage,
		PagesRead:   body.PagesRead,
		Status:      body.Status,
	})

	if err != nil {
		return err
	}

	return v.ListBySeries(c)
}

func (v *VolumeController) Delete(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	v.interactor.Delete(userId, volumeId)

	if err != nil {
		return err
	}

	return v.ListBySeries(c)
}

type volumeResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Notes       string    `json:"notes"`
	CurrentPage int       `json:"currentPage"`
	Status      string    `json:"status"`
	SeriesId    int       `json:"seriesId"`
	CreatedAt   time.Time `json:"createdAt"`
	SeriesName  string    `json:"seriesName"`
}

func volumeEntityToResponse(entity repositories.VolumeEntity) volumeResponse {
	return volumeResponse{
		Id:          entity.Id,
		Name:        entity.Name,
		CurrentPage: entity.CurrentPage,
		Notes:       entity.Notes,
		Status:      entity.Status,
		SeriesId:    entity.SeriesId,
		SeriesName:  entity.SeriesName,
	}
}

func volumeEntitiesToListResponse(volumeEntities []repositories.VolumeEntity) []volumeResponse {
	volumes := make([]volumeResponse, len(volumeEntities))

	for i, entity := range volumeEntities {
		volumes[i] = volumeEntityToResponse(entity)
	}

	return volumes
}

func (v *VolumeController) ListBySeries(c echo.Context) error {
	seriesId, err := strconv.Atoi(c.Param("seriesId"))
	if err != nil {
		return err
	}

	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	volumeEntities, err := v.interactor.ListBySeries(userId, seriesId)

	if err != nil {
		return err
	}

	volumes := volumeEntitiesToListResponse(volumeEntities)

	return c.JSON(http.StatusOK, volumes)
}

func (v *VolumeController) List(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	volumeEntities, err := v.interactor.List(userId)

	if err != nil {
		return err
	}

	volumes := volumeEntitiesToListResponse(volumeEntities)

	return c.JSON(http.StatusOK, volumes)
}

func (v *VolumeController) ListByStatus(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	status := c.Param("status")

	volumeEntities, err := v.interactor.ListByStatus(userId, status)

	if err != nil {
		return err
	}

	volumes := volumeEntitiesToListResponse(volumeEntities)

	return c.JSON(http.StatusOK, volumes)
}

func (v *VolumeController) Get(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	entity, err := v.interactor.Get(userId, volumeId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, volumeEntityToResponse(entity))
}
