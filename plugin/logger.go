package plugin

import (
	"encoding/json"
	"fmt"

	logging "github.com/ipfs/go-log"

	"github.com/ipfs/go-ipfs/plugin"
)

// LoggerPlugin controls which ipfs subsystems loggers are enabled
// These subsystems are defined in ipfs config Plugins
type LoggerPlugin struct{}

var _ plugin.Plugin = &LoggerPlugin{}

// "Plugins": {
//      "logger": {
//        "Config": {
//            "Subsystems": ["dht","relay","corerepo"],
//            "LogLevel": "error"
//        },
//        "Disabled": false
//      }
//  }
type loggerConfig struct {
	// whitelisted subsystems
	Subsystems []string

	// log level for the whitelisted subsystems
	LogLevel string
}

func (l LoggerPlugin) Name() string {
	return "logger"
}

func (l LoggerPlugin) Version() string {
	return "0.0.1"
}

// Set log levels for each subsystem
// info level for whitelisted subsystem
// fatal level for all others
func (l LoggerPlugin) Init(env *plugin.Environment) error {
	// If no plugin config given, exit with default settings
	if env == nil || env.Config == nil {
		return nil
	}

	config, err := l.loadConfig(env.Config)
	if err != nil {
		return err
	}

	// set log levels
	logging.SetAllLoggers(logging.LevelFatal)
	for _, subsystem := range config.Subsystems {
		if err := logging.SetLogLevel(subsystem, config.LogLevel); err != nil {
			return fmt.Errorf("set log level failed for subsystem: %s. Error: %s", subsystem, err.Error())
		}
	}

	return nil
}

func (l LoggerPlugin) loadConfig(envConfig interface{}) (*loggerConfig, error) {
	// load config data
	bytes, err := json.Marshal(envConfig)
	if err != nil {
		return nil, err
	}

	config := loggerConfig{
		LogLevel: "error",
	}
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
