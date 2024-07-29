package main

import (
	"go.uber.org/zap"
	router "shortener"
)

func main() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	cfg := router.NewConfig()
	router.ConfigParser("config.toml", cfg)

	r := router.SetupRouter(cfg)
	r.Run()
}
