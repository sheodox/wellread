package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sheodox/wellread/controllers"
	"github.com/sheodox/wellread/migrate"
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

	//series
	e.GET("/api/series", controllers.Series.List)
	e.POST("/api/series", controllers.Series.Add)
	e.PATCH("/api/series/:seriesId", controllers.Series.Update)
	e.DELETE("/api/series/:seriesId", controllers.Series.Delete)

	//volumes
	e.GET("/api/series/:seriesId/volumes", controllers.Volume.List)
	e.POST("/api/series/:seriesId/volumes", controllers.Volume.Add)
	e.PATCH("/api/series/:seriesId/volumes/:volumeId", controllers.Volume.Update)
	e.DELETE("/api/series/:seriesId/volumes/:volumeId", controllers.Volume.Delete)

	//reading history
	e.GET("/api/series/:seriesId/volumes/:volumeId/history", controllers.ReadingHistory.List)
	e.DELETE("/api/series/:seriesId/volumes/:volumeId/history/:historyId", controllers.ReadingHistory.Delete)

	//auth
	//e.POST("/api/auth/callback", controllers.Auth.AuthCallback)
	//e.GET("/api/auth/logout", controllers.Auth.Logout)

	// Start server
	e.Logger.Fatal(e.Start(":4004"))
}
