package config

import (
	"sync"

	"github.com/labstack/gommon/log"
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
	defaultConfig.DbName = "db_name"

	//use this if .env file (dont forget to run "source PATH_TO/.env" example "source config/.env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("your")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db_driver")
	viper.BindEnv("db_address")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	//use this for json check app.config.json for example
	// viper.AddConfigPath(".")
	// viper.SetConfigName("app.config")
	// viper.SetConfigType("json")
	// err := viper.ReadInConfig()
	// if err == nil {
	// 	fmt.Printf("Using config file: %s \n\n", viper.ConfigFileUsed())
	// }
	// finalConfig.AppPort = viper.GetString("server.port")
	// finalConfig.AppEnvironment = viper.GetString("appEnv")
	// finalConfig.DbDriver = viper.GetString("database.driver")
	// finalConfig.DbAddress = viper.GetString("database.host")
	// finalConfig.DbPort = viper.GetString("database.port")
	// finalConfig.DbUsername = viper.GetString("database.username")
	// finalConfig.DbPassword = viper.GetString("database.password")
	// finalConfig.DbName = viper.GetString("database.dbname")

	return &finalConfig
}
