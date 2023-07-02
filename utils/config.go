package utils

import (
	"github.com/spf13/viper"
	"os"
)

var (
	AppName        string
	AppHost        string
	AppPort        string
	AppEnvironment string
	AppLogLevel    string

	MysqlHost     string
	MysqlPort     string
	MysqlScheme   string
	MysqlUser     string
	MysqlPassword string
)

func InitConfig() {
	// check app environment on env
	env := os.Getenv("APP_ENV")

	// check for value
	if env == "" {
		env = "development"
	}

	if env == "development" {
		// check for config.json
		viper.SetConfigFile(`config.json`)

		// read file
		err := viper.ReadInConfig()
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// get variable for app
		AppName = viper.GetString("app.name")
		AppHost = viper.GetString("app.host")
		AppPort = viper.GetString("app.port")
		AppEnvironment = viper.GetString("app.environment")
		AppLogLevel = viper.GetString("app.log.level")

		// get variable for db
		MysqlHost = viper.GetString("database.mysql.host")
		MysqlPort = viper.GetString("database.mysql.port")
		MysqlScheme = viper.GetString("database.mysql.schema")
		MysqlUser = viper.GetString("database.mysql.user")
		MysqlPassword = viper.GetString("database.mysql.password")

		// create return
		return
	}

	if env == "staging" || env == "production" {
		// get variable for app
		AppName = os.Getenv("APP_NAME")
		AppPort = os.Getenv("APP_PORT")
		AppEnvironment = os.Getenv("APP_ENV")
		AppLogLevel = os.Getenv("APP_LOG_LEVEL")

		// get variable for db
		MysqlHost = os.Getenv("MYSQL_HOST")
		MysqlPort = os.Getenv("MYSQL_PORT")
		MysqlScheme = os.Getenv("MYSQL_SCHEME")
		MysqlUser = os.Getenv("MYSQL_USER")
		MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	}
}
