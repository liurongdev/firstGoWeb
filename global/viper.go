package global

import (
	"fmt"
	"github.com/liurongdev/firstGoWeb/middleware/redis"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

func InitViper(configName string, configPath string) *viper.Viper {
	fmt.Println("==============start init viper==============")
	Viper = viper.New()
	Viper.SetConfigName(configName) // 文件名（不带后缀）
	Viper.AddConfigPath(configPath)
	Viper.SetConfigType("yaml")
	if err := Viper.ReadInConfig(); err != nil {
		fmt.Printf("viper read config failed, err:%v\n", err)
	} else {
		fmt.Println("============init viper success=============")
		fmt.Println(Viper.AllSettings())
	}
	return Viper
}

func GetMysqlConfig() *MysqlConfig {
	config := &MysqlConfig{
		Host:     Viper.GetString("settings.database.host"),
		UserName: Viper.GetString("settings.database.username"),
		Password: Viper.GetString("settings.database.password"),
		Port:     Viper.GetInt("settings.database.port"),
		Database: Viper.GetString("settings.database.name"),
	}
	fmt.Println(config)
	return config
}

func GetRedisConfig() *redis.RedisConfig {
	return &redis.RedisConfig{
		Host:     Viper.GetString("settings.redis.host"),
		Password: Viper.GetString("settings.redis.password"),
		Port:     Viper.GetInt("settings.redis.port"),
		Database: Viper.GetString("settings.redis.database"),
	}

}
