package config

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type Config struct {
	Port       int `json:"port"`
	DbUrl      string `json:"DbUrl"`
	SomeAppKey string `json:"someAppKey"`
}

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func CreateNew() (*Config, error) {

	pr := flag.Int("port", 8080, "port")
	urlLink := flag.String("DbUrl", getEnv("DbUrl", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"), "DbUrl")
	key := flag.String("someAppKey", "testKey", "test key")
	jsonFile := flag.String("jsonFile", "", "add link to json config file")
	flag.Parse()

	var config *Config

	if *jsonFile == "" {
		config = &Config{
			*pr,
			*urlLink,
			*key,
		}
	} else {
		confJsonFile, err := ioutil.ReadFile(*jsonFile)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal(confJsonFile, &config)

	}

	_, err := url.ParseRequestURI(config.DbUrl)
	if err != nil {
		return nil,  errors.New("not URL")
	}

	/*	if strings.Contains(*url, "://") {
			config.DbUrl = *url
		} else {
			return nil, errors.New("not URL")
		}*/

	return config, nil
}
