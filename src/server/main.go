package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
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
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
	}))
	e.Use(middleware.RequestID())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	if os.Getenv("WELLREAD_ENV") == "development" {
		// allow stuff like logout redirects to route back to the dev server
		e.GET("/", func(c echo.Context) error {
			return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
		})

		// CORS for the dev server
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}))
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/usr/src/frontend",
		HTML5: true,
	}))

	//auth
	e.POST("/api/auth/callback", controllers.Auth.AuthCallback)
	e.GET("/api/auth/logout", controllers.Auth.Logout)
	e.GET("/api/auth/firebase-config", controllers.Auth.FirebaseConfig)

	authed := e.Group("/api")
	authed.Use(controllers.Auth.RequireUser)

	//series
	authed.GET("/series", controllers.Series.List)
	authed.POST("/series", controllers.Series.Add)
	authed.PATCH("/series/:seriesId", controllers.Series.Update)
	authed.DELETE("/series/:seriesId", controllers.Series.Delete)

	//volumes
	authed.GET("/series/:seriesId/volumes", controllers.Volume.List)
	authed.POST("/series/:seriesId/volumes", controllers.Volume.Add)
	authed.PATCH("/series/:seriesId/volumes/:volumeId", controllers.Volume.Update)
	authed.DELETE("/series/:seriesId/volumes/:volumeId", controllers.Volume.Delete)

	//reading history
	authed.GET("/series/:seriesId/volumes/:volumeId/history", controllers.ReadingHistory.List)
	authed.DELETE("/series/:seriesId/volumes/:volumeId/history/:historyId", controllers.ReadingHistory.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":5004"))
}
