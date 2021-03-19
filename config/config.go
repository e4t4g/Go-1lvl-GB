package config

import (
	"errors"
	"flag"
	"os"
	"strings"
)

type Config struct {
	Port         int
	Db_url       string
	Some_app_key string
}

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}
func CreateNew() (*Config, error) {

	pr := flag.Int("port", 8080, "port")
	url := flag.String("Db_url", getEnv("Db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"), "DB_URL")
	key := flag.String("some_app_key", "testkey", "test key")
	flag.Parse()
	config := &Config{
		*pr,
		*key,
		*url,
	}

	if strings.Contains(*url, "://") {
		config.Db_url = *url
	} else {
		return nil, errors.New("not URL")
	}

	return config, nil
}
