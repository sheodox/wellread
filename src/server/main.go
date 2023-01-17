package main

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
	}))
	e.Use(middleware.RequestID())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	e.GET("/health", func(c echo.Context) error {
		c.String(200, "")
		return nil
	})

	//auth
	e.POST("/api/auth/callback", controllers.Auth.AuthCallback)
	e.GET("/api/auth/logout", controllers.Auth.Logout)
	e.GET("/api/auth/firebase-config", controllers.Auth.FirebaseConfig)
	e.GET("/api/auth/status", controllers.Auth.Status)

	authed := e.Group("/api")
	authed.Use(controllers.Auth.RequireUser)

	//series
	authed.GET("/series", controllers.Series.List)
	authed.POST("/series", controllers.Series.Add)
	authed.GET("/series/:seriesId", controllers.Series.Get)
	authed.PATCH("/series/:seriesId", controllers.Series.Update)
	authed.DELETE("/series/:seriesId", controllers.Series.Delete)

	//volumes
	authed.GET("/series/:seriesId/volumes", controllers.Volume.ListBySeries)
	authed.GET("/volumes/status/:status", controllers.Volume.ListByStatus)
	authed.GET("/volumes", controllers.Volume.List)
	authed.POST("/series/:seriesId/volumes", controllers.Volume.Add)
	authed.GET("/volumes/:volumeId", controllers.Volume.Get)
	authed.PATCH("/volumes/:volumeId", controllers.Volume.Update)
	authed.DELETE("/volumes/:volumeId", controllers.Volume.Delete)

	//reading history
	authed.GET("/volumes/:volumeId/history", controllers.ReadingHistory.List)
	authed.DELETE("/history/:historyId", controllers.ReadingHistory.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":5004"))
}
