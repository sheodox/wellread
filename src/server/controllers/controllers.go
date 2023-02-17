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

func getQueryParam(c echo.Context, name string) string {
	return c.QueryParams().Get(name)
}

func getQueryParamInt(c echo.Context, name string) (int, error) {
	val := getQueryParam(c, name)

	if val == "" {
		return 0, errors.New(name + " not found in query parameters")
	}

	return strconv.Atoi(val)
}

type PageMetadata struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	TotalItems int `json:"totalItems"`
}

type PagedResults[K any] struct {
	Data K            `json:"data"`
	Page PageMetadata `json:"page"`
}
