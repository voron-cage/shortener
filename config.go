package router

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"shortener/storage"
)

func ConfigParser(configFile string, conf interface{}) {
	configBodyBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	configBody := string(configBodyBytes)

	if _, err := toml.Decode(configBody, conf); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	TarantoolAddress string                   `toml:"tarantool-address"`
	SecretKey        string                   `toml:"secret-key"`
	Tarantool        *storage.TarantoolConfig `toml:"tarantool"`
}

func NewConfig() *Config {
	return &Config{
		SecretKey: "defaultKey",
		Tarantool: storage.NewConfig(),
	}
}
