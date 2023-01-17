package controllers

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	Auth           = *NewAuthController()
	Series         = *NewSeriesController()
	Volume         = *NewVolumeController()
	ReadingHistory = *NewReadingHistoryController()
)

func getUserId(c echo.Context) (int, error) {
	userId := c.Get("UserId")

	if id, ok := userId.(int32); ok {
		return int(id), nil
	}

	return 0, errors.New("Couldn't get UserId from request")
}

func getPage(c echo.Context) (int, error) {
	page := c.QueryParams().Get("page")

	return strconv.Atoi(page)
}
