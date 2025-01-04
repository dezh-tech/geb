package userhandler

import (
	"github.com/dezh-tech/go-echo-boilerplate/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("/profile", h.userProfile,
		middleware.Auth)
}
