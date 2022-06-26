package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muharamdani/store_service/models"
)

func GetStores(c echo.Context) error {
	stores, _ := models.GetRepoStores()
	return c.JSON(http.StatusOK, stores)
}
