package userhandler

import (
	"net/http"

	"github.com/dezh-tech/go-gin-boilerplate/service/user"
	"github.com/labstack/echo/v4"
)

func (h Handler) userProfile(c echo.Context) error {
	pubkey, _ := c.Get("pubkey").(string) // not safe!

	resp, err := h.userSvc.Profile(user.ProfileRequest{Pubkey: pubkey})
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, resp)
}
