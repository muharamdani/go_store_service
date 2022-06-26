package models

import (
	"database/sql"
	"github.com/muharamdani/store_service/config"
)

type Stores struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	UserId       int            `json:"user_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Rate         int            `json:"rate"`
	TotalProduct int            `json:"total_product"`
	CreatedAt    string         `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt    string         `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func GetRepoStores() ([]Stores, error) {
	db := config.GetDBInstance()
	var stores []Stores

	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
