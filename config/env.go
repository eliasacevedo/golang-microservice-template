package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/joho/godotenv"
)

var appName = ""

func GetAppName() string {
	if appName == "" {
		appName = GetEnvVar("", true)
	}

	return appName
}

func GetPort() string {
	return GetEnvVar("_PORT", true)
}

func GetEnvFileName() string {
	return GetEnvVar("_ENV", true)
}

func GetReadTimeout() time.Duration {
	return GetTimeFromEnv("_READ_TIMEOUT", true, 60)
}

func GetReadHeaderTimeout() time.Duration {
	return GetTimeFromEnv("_READ_HEADER_TIMEOUT", true, 10)
}

func GetWriteTimeout() time.Duration {
	return GetTimeFromEnv("_WRITE_TIMEOUT", true, 60)
}

func GetIdleTimeout() time.Duration {
	return GetTimeFromEnv("_IDLE_TIMEOUT", true, 30)
}

func GetAppMode() string {
	return GetEnvVar("_MODE", true)
}

func GetDescription() string {
	return GetEnvVar("_DESCRIPTION", false)
}

func GetVersion() int64 {
	number, _ := GetNumberFromEnv("_VERSION", false, 1)
	return number
}

func GetRoutesPrefix() string {
	return GetStringFromEnv("ROUTES_PREFIX", true, "api")
}

func GetTimeBeforeShutdownServer() time.Duration {
	return GetTimeFromEnv("_TIME_BEFORE_SHUTDOWN_SERVER", false, 60)
}

func GetMustLogInfo() bool {
	return GetBooleanFromEnv("_LOG_INFO", false, true)
}

func GetMustLogServerError() bool {
	return GetBooleanFromEnv("_LOG_SERVER_ERRORS", false, true)
}

func GetMustLogValidationError() bool {
	return GetBooleanFromEnv("_LOG_VALIDATION_ERROR", false, true)
}

func GetMustLogHTTPBeginRequestInfo() bool {
	return GetBooleanFromEnv("_LOG_HTTP_BEGIN_REQUEST", false, true)
}

func GetMustLogHTTPError() bool {
	return GetBooleanFromEnv("_LOG_HTTP_ERROR", false, true)
}

func GetMustLogHTTPEndRequestInfo() bool {
	return GetBooleanFromEnv("_LOG_HTTP_END_REQUEST", false, true)
}

func GetDatabase() string {
	return GetEnvVar("_CONNECTION_STRING", true)
}

func GetStringFromEnv(key string, required bool, initial string) string {
	prefix := ""
	if key == "" {
		prefix = "APPNAME"
	} else {
		prefix = GetAppName()
	}
	index := prefix + key
	value, isThere := os.LookupEnv(index)
	if !isThere && required {
		return initial
	}
	return value
}

func GetTimeFromEnv(key string, required bool, initial time.Duration) time.Duration {
	number, found := GetNumberFromEnv(key, false, 0)
	if !found {
		return initial
	}
	return time.Duration(number)
}

func GetNumberFromEnv(key string, required bool, initial int64) (int64, bool) {
	value := GetEnvVar(key, required)
	if value == "" && !required {
		return initial, false
	}

	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(key + " param is not a valid number")
	}

	return num, true
}

func GetBooleanFromEnv(key string, required bool, initial bool) bool {
	value := GetEnvVar(key, required)
	if value == "" && !required {
		return initial
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		panic(key + " param is not a valid number")
	}
	return result
}

func GetEnvVar(key string, required bool) string {
	prefix := ""
	if key == "" {
		prefix = "APPNAME"
	} else {
		prefix = GetAppName()
	}
	index := prefix + key
	value, isThere := os.LookupEnv(index)
	if !isThere && required {
		panic(fmt.Sprintf("%s env var is not defined", index))
	}
	return value
}

func LoadEnvFromFile(l utilities.Logger) {
	environment := os.Getenv("env")
	if environment == "" {
		environment = ".env.local"
	}

	err := godotenv.Load(environment)
	if err != nil {
		l.PanicApp(fmt.Sprintf("error loading env file: %s", err.Error()))
	}
}
