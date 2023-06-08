package main

import (
	"nunuim/cmd/migration/wire"
	"nunuim/pkg/config"
	"nunuim/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}
