package controllers

import (
	"errors"

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

	if id, ok := userId.(int); ok {
		return id, nil
	}

	return 0, errors.New("Couldn't get UserId from request")
}
