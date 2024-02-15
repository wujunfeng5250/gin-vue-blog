package core

import (
	"fmt"
	"gin-vue-blog/server/config"
	"gin-vue-blog/server/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
