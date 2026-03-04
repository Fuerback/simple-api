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
		UnsafeAllowOriginFunc: func(c *echo.Context, origin string) (string, bool, error) {
			return origin, true, nil
		},
		AllowMethods:     []string{http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	e.POST("/api/pixel/events", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Render requires binding to 0.0.0.0; PORT defaults to 10000 on Render (see https://render.com/docs/web-services#port-binding)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := "0.0.0.0:" + port
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
