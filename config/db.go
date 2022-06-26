package config

import "fmt"

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
