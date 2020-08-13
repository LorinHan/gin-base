package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
		MaxSize    int    `yaml:"maxSize"`
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

/**
 * @description: 读取配置文件 "config.yml"
 * 	配置文件一般在项目根目录，运行入口程序的话，os.Getwd()加上文件名即可得到配置文件
 *	但是在某些情况下，例如单元测试的时候，程序在某个包下运行，os.Getwd()获取到的对应包的路径而不是项目的根路径
 *	遍历程序运行的路径，得到正确的配置文件路径后进行读取并解析配置类
 * @author: Lorin
 * @time: 2020/8/13 下午6:19
 */
func init() {
	pwd, _ := os.Getwd()
	sp := string(os.PathSeparator)
	splits := strings.Split(pwd, sp)
	for i := len(splits); i > 0; i-- {
		path := strings.Join(splits[0:i], sp) + sp + "config.yml"
		if pathExists(path) {
			configFile, fileErr := ioutil.ReadFile(path)
			if fileErr != nil {
				log.Print(fileErr)
			}
			err := yaml.Unmarshal(configFile, &Conf)
			if err != nil {
				log.Print("config file parse error")
			}
			break
		}
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
