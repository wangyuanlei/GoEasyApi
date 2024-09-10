package libraries

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Sqlite string `yaml:"sqlite"`
	} `yaml:"config"`
	IpFilterMode string `yaml:"ipFilterMode"`
	Cache        bool   `yaml:"cache"`
	Redis        struct {
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
}

var Conf Config

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func WriteConfig() {
	yamlData, err := yaml.Marshal(&Conf)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}
	err = ioutil.WriteFile("config.yml", yamlData, 0644)
	if err != nil {
		log.Fatalf("WriteFile: %v", err)
	}
}
