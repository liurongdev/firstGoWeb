package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysql(mysqlConfig *MysqlConfig) {
	fmt.Println("==============start init mysql==============")
	dsn := "root:root@tcp(127.0.0.1:3306)/ferry?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open mysql err: %v", err)
	}
	fmt.Println("==============init mysql success==============")
}

func GetMysql() *gorm.DB {
	return DB
}

type MysqlConfig struct {
	Host     string
	UserName string
	Password string
	Port     int
	Database string
}
