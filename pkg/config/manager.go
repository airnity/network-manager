package config

import (
	log "github.com/sirupsen/logrus"
)

// Manager.
//go:generate mockgen -destination=./mocks/mock_Manager.go -package=mocks easymile.com/cloud/vehicle-services/pkg/vehicle-services/config Manager
type Manager interface {
	// Load configuration
	Load() error
	// Get configuration object
	GetConfig() *Config
	// // Add on change hook for configuration change
	// AddOnChangeHook(hook func())
}

func NewManager(logger *log.Logger) Manager {
	return &configmanager{logger: logger}
}
