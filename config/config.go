package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type ConfigIPForwarding struct {
	Enabled bool   `mapstructure:"enabled"`
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
}

//AppConfig Application configuration
type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbAddress      string `mapstructure:"db_address"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	var finalConfig AppConfig

	defaultConfig.AppPort = 8080
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbAddress = "localhost"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "root"
	defaultConfig.DbPassword = "1"
	defaultConfig.DbName = "alta_final"

	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("app.config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n\n", viper.ConfigFileUsed())
	}
	finalConfig.AppPort = viper.GetInt("server.port")
	finalConfig.AppEnvironment = viper.GetString("appEnv")
	finalConfig.DbDriver = viper.GetString("database.driver")
	finalConfig.DbAddress = viper.GetString("database.host")
	finalConfig.DbPort = viper.GetInt("database.port")
	finalConfig.DbUsername = viper.GetString("database.username")
	finalConfig.DbPassword = viper.GetString("database.password")
	finalConfig.DbName = viper.GetString("database.dbname")

	return &finalConfig
}
