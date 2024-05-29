package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

func InitDB() *gorm.DB {
	viper.SetConfigFile("prod.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file :", err)
		return nil
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config :", err)
		return nil
	}

	host := config.Database.Host
	fmt.Println(host)
	port := config.Database.Port
	fmt.Println(port)
	dbname := config.Database.DBName
	fmt.Println(dbname)
	username := config.Database.Username
	fmt.Println(username)
	password := config.Database.Password
	fmt.Println(password)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})

	if err != nil {
		panic("Can't connect to database")
	}
	return db
}
