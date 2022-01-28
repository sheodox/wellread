package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
)

type ReadingHistoryController struct {
	interactor *interactors.ReadingHistoryInteractor
}

func NewReadingHistoryController() *ReadingHistoryController {
	return &ReadingHistoryController{&interactors.ReadingHistory}
}

type readingHistoryResponse struct {
	Id          int       `json:"id"`
	CurrentPage int       `json:"currentPage"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (v *ReadingHistoryController) Delete(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("historyId"))
	if err != nil {
		return err
	}

	v.interactor.Delete(volumeId)

	if err != nil {
		return err
	}

	return v.List(c)
}

func (v *ReadingHistoryController) List(c echo.Context) error {
	volumeId, err := strconv.Atoi(c.Param("volumeId"))
	if err != nil {
		return err
	}

	historyEntities, err := v.interactor.List(volumeId)

	if err != nil {
		return err
	}

	historyResponse := make([]readingHistoryResponse, len(historyEntities))

	for i, history := range historyEntities {
		historyResponse[i] = readingHistoryResponse{
			Id:          history.Id,
			CurrentPage: history.CurrentPage,
			CreatedAt:   history.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, historyResponse)
}
