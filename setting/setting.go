package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	RunMode  string   `yaml:"runMode"`
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
	LogPath  string   `yaml:"logPath"`
}

type server struct {
	HttpPort int `yaml:"HttpPort"`
	ReadTimeout time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type database struct {
	Type string `yaml:"type"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	DBName string `yaml:"dbname"`
	Network string `yaml:"network"`
	Port int `yaml:"port"`
}

type envType string
const (
	ENV_TYPE_LOCAL envType = "local"
	ENV_TYPE_DEV   envType = "develop"
	ENV_TYPE_PRO   envType = "product"
)

// 当前环境
var currentEnv envType = ENV_TYPE_DEV

var config = &Config{}

func InitSetting() {
	yamlFile, err := ioutil.ReadFile("conf/" + string(currentEnv) + ".yaml")
	if err != nil {
		log.Fatalln("read config failed:", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalln("yaml file unmarshal failed:", err)
		return
	}
}

func Setting() *Config {
	return config
}
