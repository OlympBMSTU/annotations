package config

import (
	"reflect"
	"fmt"
	"strings"
	"errors"
	"os"
	"io/ioutil"
)

type Config struct {
	fileStorageDir string
	listenerHost   string
	listenerPort   string
	dbHost         string
	dbPort         string
	database       string
	dbUser         string
	dbPassword     string
	testVersion    string
	authCookieName string
	authSecret     string
}

var configInstance *Config = nil

func (cfg Config) GetFileStorageName() string {
	return cfg.fileStorageDir
}

func (cfg Config) GetDBHost() string {
	return cfg.dbHost
}

func (cfg Config) GetDBPort() string {
	return cfg.dbPort
}

func (cfg *Config) GetDatabase() string {
	return cfg.database
}

func (cfg Config) GetDBUser() string {
	return cfg.dbUser
}

func (cfg Config) GetDBPassword() string {
	return cfg.dbPassword
}

func (cfg Config) IsTest() bool {
	return cfg.testVersion == "test"
}

func (cfg Config) GetAuthCookieName() string {
	return cfg.authCookieName
}

func (cfg Config) GetAuthSecret() string {
	return cfg.authSecret
}

func (cfg Config) GetListenerHost() string {
	return cfg.listenerHost
}

func (cfg Config) GetListenerPort() string {
	return cfg.listenerPort
}

func initConfig() error {
	iniPath := "/etc/annotations.conf" //"/etc/serv"

	file, err := os.Open(iniPath)

	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Cant start server without initial file\n" +
			"Please creaate init file and put it to /etc/serv/")
		return err
	}

	fileData := string(fbytes)
	configs := strings.Split(fileData, "\n")

	countFields := reflect.ValueOf(Config{}).NumField()
	if countFields > len(configs) {
		return errors.New("Incorrect count fields in config")
	}


	configInstance = &Config{
		configs[0],
		configs[1],
		configs[2],
		configs[3],
		configs[4],
		configs[5],
		configs[6],
		configs[7],
		configs[8],
		configs[9],
		configs[10],
	}
	
	return nil
}

func GetConfigInstance() (*Config, error) {
	if configInstance == nil {
		err := initConfig()
		if err != nil {
			return nil, err
		}
	}
	return configInstance, nil
}
