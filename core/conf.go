package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

var Conf *Config

func InitConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取配置文件失败: %v", err))
	}

	Conf = &Config{}
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("解析配置文件失败: %v", err))
	}

	fmt.Println("配置文件加载成功!")
}
