package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muharamdani/store_service/config"
	"github.com/muharamdani/store_service/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	config.NewDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/stores", controllers.GetStores)

	return e
}
