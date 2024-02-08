package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsers(c echo.Context) error {
	u, err := h.userStore.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, u)
}
