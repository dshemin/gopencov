package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

// Run run application REST API server
func Run(ctx context.Context, logger echo.Logger, address string) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Logger = logger

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return nil
	})

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(shutdownCtx); err != nil {
			e.Logger.Errorf("cannot stop server: %s", err)
		}
	}()

	err := e.Start(address)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
