package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

/**
    conf
    @author: roccoshi
    @desc: configuration
**/

type Conf struct {
	DB      MySQL   `yaml:"mysql"`
	Session Session `yaml:"session"`
}

type MySQL struct {
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Port        int    `yaml:"port"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	MaxLifeTime int    `yaml:"max_life_time"`
}

type Session struct {
	Secret string `yaml:"secret"`
	Name   string `yaml:"name"`
}

var Config *Conf

func LoadConfig() {
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
}
