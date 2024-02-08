package main

import (
	"net/http"
	"os"

	"echo.osaxon/database"
	"echo.osaxon/handler"
	"echo.osaxon/store"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db, err := database.New()
	database.Migrate(db)

	if err != nil {
		app.Logger.Fatal(err)
	}

	port := os.Getenv("PORT")

	us := store.NewUserStore(db)

	h := handler.NewHandler(us)

	h.SetupRoutes(app.Group("/v1"))

	app.Logger.Fatal(app.Start("0.0.0.0:" + port))
}
