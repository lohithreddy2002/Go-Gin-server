package config

import "fmt"

const (
	DBUser     = "root"
	DBPassword = "lohith"
	DBName     = "mydb"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBtype     = "mysql"
)

func GetDBType() string {
	return DBtype
}

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}
