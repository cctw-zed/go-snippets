package designpattern

import (
	"sync"
)

type Config struct {
	DatabaseURL string
	APIKey      string
}

var configInstance *Config
var once sync.Once

func GetConfigInstance() *Config {
	once.Do(func() {
		configInstance = &Config{
			DatabaseURL: "localhost:3306",
			APIKey:      "abc123",
		}
	})
	return configInstance
}
