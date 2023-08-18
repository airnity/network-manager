package vrf

import (
	"airnity.com/router-sidecar/pkg/config"

	log "github.com/sirupsen/logrus"
)

type Client interface {
	Synchronize() error
}

func NewClient(cfgManager config.Manager, logger *log.Logger) Client {
	return &client{cfg: cfgManager, logger: logger}
}
