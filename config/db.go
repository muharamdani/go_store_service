package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func GetDBType() string {
	return "mysql"
}

func GetMySQLConnectionString() string {
	conf := GetConfig()
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DbUsername,
		conf.DbPass,
		conf.DbHost,
		conf.DbPort,
		conf.DbName)

	return dataBase
}

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
