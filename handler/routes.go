package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) SetupRoutes(v1 *echo.Group) {
	v1.GET("/hello", h.HelloWorld)
}
