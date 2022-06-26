package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muharamdani/store_service/model"
	"github.com/muharamdani/store_service/storage"
)

func GetStores(c echo.Context) error {
	stores, _ := GetRepoStores()
	return c.JSON(http.StatusOK, stores)
}

func GetRepoStores() ([]model.Stores, error) {
	db := storage.GetDBInstance()
	stores := []model.Stores{}

	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
