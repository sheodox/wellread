package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sheodox/wellread/interactors"
)

type AuthController struct {
	interactor *interactors.AuthInteractor
}

func NewAuthController() *AuthController {
	return &AuthController{
		&interactors.Auth,
	}
}

type authRequest struct {
	IdToken string `json:"idToken"`
}

type authResponse struct {
}

func (a *AuthController) AuthCallback(c echo.Context) error {

	authBody := new(authRequest)
	if err := c.Bind(authBody); err != nil {
		return err
	}

	user, err := a.interactor.Login(authBody.IdToken)

	if err != nil {
		return err
	}

	// authenticate the user
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}

	sess.Values["wellread_user_id"] = user.Id
	err = sess.Save(c.Request(), c.Response())

	if err != nil {
		return err
	}

	response := authResponse{}

	return c.JSON(http.StatusOK, response)
}

func (a *AuthController) RequireUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)

		if err != nil {
			return err
		}

		val := sess.Values["wellread_user_id"]

		if userId, ok := val.(int); ok {
			user, err := a.interactor.GetUser(userId)

			if err != nil {
				return err
			}

			c.Set("User", user)
			c.Set("UserId", user.Id)

			return next(c)
		} else {
			c.String(http.StatusUnauthorized, "")
		}
		return nil
	}
}

func (a *AuthController) Logout(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}

	sess.Values["wellread_user_id"] = nil
	err = sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (a *AuthController) FirebaseConfig(c echo.Context) error {
	return c.File("./public-firebase-config.json")
}
