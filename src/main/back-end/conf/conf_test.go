package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"testing"
)

/**
    conf
    @author: roccoshi
    @desc: test
**/

func TestLoadConfig(t *testing.T) {
	file, err := ioutil.ReadFile("../config.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
	t.Logf("%#v", Config)
	// 1 - mysql config
	dbConf := Config.DB
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/eth_wallet?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port)
	t.Logf("%s", dsn)
}
