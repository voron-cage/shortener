package storage

import (
	"shortener/common"
)

type TarantoolConfig struct {
	ListenAddress   string           `toml:"listen-address"`
	User            string           `toml:"user"`
	ConnectTimeout  *common.Duration `toml:"connect-timeout"`
	ResponseTimeout *common.Duration `toml:"response-timeout"`
}

func NewConfig() *TarantoolConfig {
	return &TarantoolConfig{
		ListenAddress: "127.0.0.1:3301",
	}
}
