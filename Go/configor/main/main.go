package main

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}

func main() {
	var config1 = Config{}
	var config2 = Config{}
	var config3 = Config{}
	var config4 = Config{}

	configor.Load(&config1, "config.yml")
	configor.Load(&config2, "config.toml")
	configor.Load(&config3, "config.json")
	configor.Load(&config4, "config.ini")
	fmt.Printf("yaml config: %#v \n\n", config1)
	fmt.Printf("toml config: %#v \n\n", config2)
	fmt.Printf("json config: %#v \n\n", config3)
	fmt.Printf("ini config: %#v \n\n", config4)
}
