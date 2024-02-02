package main

import (
	"net/http"

	"echo.osaxon/database"
	"echo.osaxon/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	_, err := database.Connect()

	if err != nil {
		app.Logger.Fatal(err)
	}

	h := handler.NewHandler()
	h.SetupRoutes(app.Group("/v1"))

	app.Logger.Fatal(app.Start(":1323"))
}
