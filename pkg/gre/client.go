package gre

import (
	"airnity.com/network-manager/pkg/config"

	log "github.com/sirupsen/logrus"
)

type Client interface {
	Synchronize() error
}

func NewClient(cfgManager config.Manager, logger *log.Logger) Client {
	return &client{cfg: cfgManager, logger: logger}
}
