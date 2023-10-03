package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{Message: "Klk"})
	})
	e.GET("/api/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Current string `json:"current"`
		}{Current: time.Now().Format(time.RFC3339)})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
