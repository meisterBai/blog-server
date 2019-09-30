package config

import (
	"server/logger"
	"github.com/pelletier/go-toml"
)


var (
	Conf = New()
)

func New() *toml.Tree {
	config, err := toml.LoadFile("./config/config.toml")
	if err != nil {
		logger.Error(err)
	}

	return config
}

