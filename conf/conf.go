package conf

import (
	"gin-base/utils/pathUtil"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Swagger bool `yaml:"swagger"`
	}
	Database struct {
		DbType       string `yaml:"dbType"`
		LogMode      bool   `yaml:"logMode"`
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
		MaxSize    int    `yaml:"maxSize"`
		LogFormat  string `yaml:"logFormat"`
		ToStd      bool   `yaml:"toStd"`
	}
	Jwt struct {
		Expired time.Duration `yaml:"expired"`
		Secret  string        `yaml:"secret"`
	}
	PageSize uint `yaml:"pageSize"`
}

var (
	Conf     = &Config{}
	Jwt      = &Conf.Jwt
	Database = &Conf.Database
	LogConf  = &Conf.Log
	Server   = &Conf.Server
	PageSize = &Conf.PageSize
)

/**
 * @description: 读取配置文件 "config.yml"
 * 	配置文件一般在项目根目录，运行入口程序的话，os.Getwd()加上文件名即可得到配置文件
 *	但是在某些情况下，例如单元测试的时候，程序在某个包下运行，os.Getwd()获取到的对应包的路径而不是项目的根路径
 *	遍历程序运行的路径，得到正确的配置文件路径后进行读取并解析配置类
 * @author: Lorin
 * @time: 2020/8/13 下午6:19
 */
func init() {
	path := pathUtil.FilePath("config.yml")
	if path == "" {
		log.Fatal("could not found config file")
	}
	configFile, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		log.Print(fileErr)
	}
	err := yaml.Unmarshal(configFile, &Conf)
	if err != nil {
		log.Print("config file parse error")
	}
}
