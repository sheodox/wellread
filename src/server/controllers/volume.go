package controllers

import (
	"database/sql"
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

	entity, err := v.interactor.Add(userId, seriesId, body.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, volumeResponse{
		Id:          int(entity.ID),
		Name:        entity.Name,
		CurrentPage: int(entity.CurrentPage),
		Notes:       entity.Notes,
		Status:      entity.Status,
		SeriesId:    int(entity.SeriesID),
	})
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

	return c.NoContent(http.StatusOK)
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

	return c.NoContent(http.StatusOK)
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

func (v *VolumeController) List(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	seriesId, seriesIdErr := getQueryParamInt(c, "seriesId")
	name := getQueryParam(c, "name")
	status := getQueryParam(c, "status")
	page, pageErr := getQueryParamInt(c, "page")

	// if we actually got a page we get no error, change the page to start at 0, not 1 like the UI finds handy
	if pageErr == nil {
		page -= 1
	}

	volumeEntities, err := v.interactor.List(
		userId,
		sql.NullInt32{
			Int32: int32(seriesId),
			Valid: seriesIdErr == nil,
		},
		sql.NullString{
			String: name,
			Valid:  name != "",
		},
		sql.NullString{String: status, Valid: status != ""},
		page)

	if err != nil {
		return err
	}

	totalItems := 0
	if len(volumeEntities) > 0 {
		totalItems = int(volumeEntities[0].TotalResults)
	}

	volumes := make([]volumeResponse, len(volumeEntities))

	for i, entity := range volumeEntities {
		volumes[i] = volumeResponse{
			Id:          int(entity.ID),
			Name:        entity.Name,
			CurrentPage: int(entity.CurrentPage),
			Notes:       entity.Notes,
			Status:      entity.Status,
			SeriesId:    int(entity.SeriesID),
			SeriesName:  entity.SeriesName,
		}
	}

	return c.JSON(http.StatusOK, PagedResults[[]volumeResponse]{
		Data: volumes,
		Page: PageMetadata{
			// change the page back to starting at 1 for the UI
			PageNumber: page + 1,
			TotalItems: totalItems,
			PageSize:   repositories.PageSize,
		},
	})
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

	return c.JSON(http.StatusOK, volumeResponse{
		Id:          int(entity.ID),
		Name:        entity.Name,
		CurrentPage: int(entity.CurrentPage),
		Notes:       entity.Notes,
		Status:      entity.Status,
		SeriesId:    int(entity.SeriesID),
		SeriesName:  entity.SeriesName,
	})
}
