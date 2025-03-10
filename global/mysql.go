package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ferry?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open mysql err: %v", err)
	}
}
func GetMysql() *gorm.DB {
	return DB
}
