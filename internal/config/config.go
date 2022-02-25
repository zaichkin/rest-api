package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	SrvHost      string `json:"ServerHost"`
	SrvPort      string `json:"ServerPort"`
	PgUser       string `json:"PgUser"`
	PgPasswd     string `json:"PgPasswd"`
	PgHost       string `json:"PgHost"`
	PgPort       string `json:"PgPort"`
	PgDB         string `json:"PgDB"`
	CountConnect int    `json:"CountConnect"`
}

func GetConfig(pathConfig string) *Config {
	var conf *Config
	file, err := os.Open(pathConfig)
	if err != nil {
		fmt.Println("Can't open file ", err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Can't read file ", err)
	}

	if err = json.Unmarshal(buffer, &conf); err != nil {
		fmt.Println("Can't unmarshal file ", err)
	}

	return conf
}
