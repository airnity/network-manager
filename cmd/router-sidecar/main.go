package main

import (
	"airnity.com/router-sidecar/pkg/config"
	"airnity.com/router-sidecar/pkg/nat"
	"airnity.com/router-sidecar/pkg/vrf"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Create new logger
	logger := log.New()

	// Create configuration manager
	cfgManager := config.NewManager(logger)

	// Load configuration
	err := cfgManager.Load()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	logCfg := cfgManager.GetConfig().Logs
	if logCfg != nil {
		switch logCfg.Level {
		case "debug":
			logger.SetLevel(log.DebugLevel)
		case "error":
			logger.SetLevel(log.ErrorLevel)
		case "fatal":
			logger.SetLevel(log.FatalLevel)
		case "warn":
			logger.SetLevel(log.WarnLevel)
		case "panic":
			logger.SetLevel(log.PanicLevel)
		case "trace":
			logger.SetLevel(log.TraceLevel)
		default:
			logger.SetLevel(log.InfoLevel)
		}

		switch logCfg.Format {
		case "json":
			logger.SetFormatter(&log.JSONFormatter{})
		default:
			logger.SetFormatter(&log.TextFormatter{})
		}
	}
	vrfClient := vrf.NewClient(cfgManager, logger)
	natClient := nat.NewClient(cfgManager, logger)

	vrfClient.Synchronize()
	natClient.Synchronize()
}
