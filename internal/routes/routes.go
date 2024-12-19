package routes

import (
	_ "database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "log"
)

const dsn = "root@tcp(127.0.0.1:3306)/kleaners_0"

func InitializeDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
