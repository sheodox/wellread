package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sheodox/bookmark/db"
	"github.com/sheodox/bookmark/migrate"
	"github.com/sheodox/bookmark/services/series"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	err = migrate.RunMigrations()
	if err != nil {
		log.Fatal("Error running migrations ", err)
	}

	dbConn, err := db.Connect()

	if err != nil {
		log.Fatal("Error connecting to database ", err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/usr/src/frontend",
		HTML5: true,
	}))

	// Routes
	e.GET("/", hello)

	seriesService := series.New(dbConn)

	e.GET("/api/series", seriesService.List)
	e.POST("/api/series", seriesService.Add)
	e.PATCH("/api/series/:seriesId", seriesService.Update)
	e.DELETE("/api/series/:seriesId", seriesService.Delete)
	e.GET("/api/series/:seriesId/volumes", seriesService.ListVolumes)
	e.POST("/api/series/:seriesId/volumes", seriesService.AddVolume)
	e.PATCH("/api/series/:seriesId/volumes/:volumeId", seriesService.UpdateVolume)
	e.DELETE("/api/series/:seriesId/volumes/:volumeId", seriesService.DeleteVolume)

	// Start server
	e.Logger.Fatal(e.Start(":4004"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
