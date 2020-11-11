package environment

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var Port int
var DbHost string
var DbPort int
var DbPassword string
var DbName string
var DbUser string
var LogLevel string

func init() {
	LoadConfigs()
}

func LoadConfigs() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)

	var err error

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Error(err.Error())
		Port = 8081
	}

	DbHost = os.Getenv("DB_HOST")

	DbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error(err.Error())
		DbPort = 5432
	}

	DbPassword = os.Getenv("DB_PASSWORD")

	DbName = os.Getenv("DB_NAME")

	DbUser = os.Getenv("DB_USER")

	LogLevel = os.Getenv("LOG_LEVEL")
	if LogLevel == "" {
		LogLevel = "info"
	}

}

func GetLogLevel() log.Level {
	switch LogLevel {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	default:
		return log.ErrorLevel
	}
}