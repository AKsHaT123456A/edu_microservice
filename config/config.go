package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
