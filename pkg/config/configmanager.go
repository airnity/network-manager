package config

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// Main configuration folder path.
var mainConfigFolderPath = "/etc/router-sidecar/conf"

var validate = validator.New()

type configmanager struct {
	cfg     *Config
	configs []*viper.Viper
	logger  *log.Logger
}

func (ctx *configmanager) Load() error {
	// List files
	files, err := ioutil.ReadDir(mainConfigFolderPath)
	if err != nil {
		return err
	}

	// Generate viper instances for static configs
	ctx.configs = generateViperInstances(files)

	// Load configuration
	err = ctx.loadConfiguration()
	if err != nil {
		return err
	}

	return nil
}

// GetConfig allow to get configuration object.
func (ctx *configmanager) GetConfig() *Config {
	return ctx.cfg
}

func (ctx *configmanager) loadConfiguration() error {
	// Create a viper instance for default value and merging
	globalViper := viper.New()

	// Loop over configs
	for _, vip := range ctx.configs {
		// Read configuration
		err := vip.ReadInConfig()
		// Check error
		if err != nil {
			return err
		}

		// Merge all configurations
		err = globalViper.MergeConfigMap(vip.AllSettings())
		// Check error
		if err != nil {
			return err
		}
	}

	// Prepare configuration object
	var out Config
	// Quick unmarshal.
	err := globalViper.Unmarshal(&out)
	// Check error
	if err != nil {
		return err
	}

	// Configuration validation
	err = validate.Struct(out)
	// Check error
	if err != nil {
		return err
	}

	ctx.cfg = &out

	return nil
}

func generateViperInstances(files []os.FileInfo) []*viper.Viper {
	list := make([]*viper.Viper, 0)
	// Loop over static files to create viper instance for them
	funk.ForEach(files, func(file os.FileInfo) {
		filename := file.Name()
		// Create config file name
		cfgFileName := strings.TrimSuffix(filename, path.Ext(filename))
		// Test if config file name is compliant (ignore hidden files like .keep or directory)
		if !strings.HasPrefix(filename, ".") && cfgFileName != "" && !file.IsDir() {
			// Create new viper instance
			vip := viper.New()
			// Set config name
			vip.SetConfigName(cfgFileName)
			// Add configuration path
			vip.AddConfigPath(mainConfigFolderPath)
			// Append it
			list = append(list, vip)
		}
	})
	return list
}
