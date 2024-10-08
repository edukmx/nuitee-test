package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type Config struct {
	Port              string
	ClientTimeout     time.Duration
	HotelBedsConfig   HotelBedsConfig
	NationalitiesPath string
}

type HotelBedsConfig struct {
	Host   string
	ApiKey string
	Secret string
}

func NewConfig() *Config {
	timeoutStr := os.Getenv("CLIENT_TIMEOUT")
	timeoutInt, err := strconv.Atoi(timeoutStr)
	if err != nil {
		slog.Error(fmt.Sprintf("Error converting timeout to int: %s", err))
		timeoutInt = 5
	}

	timeout := time.Duration(timeoutInt) * time.Second

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	csvFilePath := filepath.Join(basePath, "../../data/nationalities.csv")

	return &Config{
		Port:          os.Getenv("SERVER_PORT"),
		ClientTimeout: timeout,
		HotelBedsConfig: HotelBedsConfig{
			Host:   os.Getenv("HOTELBEDS_HOST"),
			ApiKey: os.Getenv("HOTELBEDS_API_KEY"),
			Secret: os.Getenv("HOTELBEDS_SECRET"),
		},
		NationalitiesPath: csvFilePath,
	}
}
