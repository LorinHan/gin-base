package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Database struct {
		DbType string `yaml:"dbtype"`
		Host string `yaml:"host"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		DbName string `yaml:"dbname"`
		Args string `yaml:"args"`
	}
	LogLevel string `yaml:"loglevel"`
}

var Conf = &Config{}

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
