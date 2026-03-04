package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{"Content-Type"},
	}))

	e.POST("/api/pixel/events", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
