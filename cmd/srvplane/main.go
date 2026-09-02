package main

import (
	"embed"
	"io/fs"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

//go:embed dist/*
var embedFS embed.FS

func main() {
	distFS, err := fs.Sub(embedFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/ping", func(c *echo.Context) error {
		return c.String(200, "pong")
	})
	apiV1.Any("*", func(c *echo.Context) error {
		return c.String(404, "not found")
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       ".",
		Filesystem: distFS,
		HTML5:      true,
	}))

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
