package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	Database struct {
		DbType       string `yaml:"dbType"`
		Host         string `yaml:"host"`
		UserName     string `yaml:"userName"`
		Password     string `yaml:"passWord"`
		DbName       string `yaml:"dbName"`
		Args         string `yaml:"args"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxLifetime  int    `yaml:"maxLifetime"`
	}
	Log struct {
		Level      string `yaml:"level"`
		Path       string `yaml:"path"`
		MaxBackups int    `yaml:"maxBackups"`
		MaxAge     int    `yaml:"maxAge"`
		MaxSize    int    `yaml:maxSize`
		LogFormat  string `yaml:"logFormat"`
		ToStd      bool   `yaml:"toStd"`
	}
	Jwt struct {
		Expired time.Duration `yaml:"expired"`
		Secret  string        `yaml:"secret"`
	}
}

var (
	Conf     = &Config{}
	Jwt      = &Conf.Jwt
	Database = &Conf.Database
	LogConf  = &Conf.Log
)

func init() {
	pwd, _ := os.Getwd()
	configFile, fileErr := ioutil.ReadFile(pwd + "/config.yml")
	if fileErr != nil {
		log.Print(fileErr)
	}
	err := yaml.Unmarshal(configFile, &Conf)
	if err != nil {
		log.Print("config file error")
	}
}
