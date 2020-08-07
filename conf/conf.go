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
		DbType string `yaml:"dbtype"`
		Host string `yaml:"host"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		DbName string `yaml:"dbname"`
		Args string `yaml:"args"`
		MaxIdleConns int `yaml:"maxIdleConns"`
		MaxOpenConns int `yaml:"maxOpenConns"`
		MaxLifetime int `yaml:"maxLifetime"`
	}
	LogLevel string `yaml:"loglevel"`
	Jwt struct {
		Expired time.Duration `yaml:"expired"`
		Secret string `yaml:"secret"`
	}
}

var (
	Conf = &Config{}
	Jwt = &Conf.Jwt
	Database = &Conf.Database
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
