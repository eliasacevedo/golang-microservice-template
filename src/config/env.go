package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/joho/godotenv"
)

func GetAppName() string {
	return GetEnvVar("", true)
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

func GetTimeBeforeShutdownServer() time.Duration {
	return GetTimeFromEnv("_TIME_BEFORE_SHUTDOWN_SERVER", false, 60)
}

func GetTimeFromEnv(key string, required bool, initial time.Duration) time.Duration {
	value := GetEnvVar(key, required)
	if value == "" && !required {
		return initial
	}

	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(key + " param is not a valid number")
	}
	return time.Duration(num)
}

func GetEnvVar(key string, required bool) string {
	prefix := ""
	if key == "" {
		prefix = "APPNAME"
	} else {
		prefix = GetAppName()
	}

	value, isThere := os.LookupEnv(prefix + key)
	if !isThere && required {
		panic(fmt.Sprintf("%s env var is not defined ", key))
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
