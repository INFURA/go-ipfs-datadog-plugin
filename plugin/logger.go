package plugin

import (
	"encoding/json"
	"fmt"

	logging "github.com/ipfs/go-log"

	"github.com/ipfs/go-ipfs/plugin"
)

const defaultLevel = "error"

var _ plugin.Plugin = &LoggerPlugin{}

type loggerConfig struct {
	Levels       map[string][]string
	DefaultLevel string
}

type LoggerPlugin struct{}

func (l LoggerPlugin) Name() string {
	return "datadog-logger"
}

func (l LoggerPlugin) Version() string {
	return "0.0.1"
}

// Set log levels for each system (logger)
func (l LoggerPlugin) Init(env *plugin.Environment) error {
	// If no plugin config given, exit with default settings
	if env == nil || env.Config == nil {
		return nil
	}

	config, err := l.loadConfig(env.Config)
	if err != nil {
		return err
	}

	// set default log level for all loggers
	defaultLevel, err := logging.LevelFromString(config.DefaultLevel)
	if err != nil {
		return err
	}

	logging.SetAllLoggers(defaultLevel)
	for level, subsystems := range config.Levels {
		for _, subsystem := range subsystems {
			if err := logging.SetLogLevel(subsystem, level); err != nil {
				return fmt.Errorf("set log level failed for subsystem: %s. Error: %s", subsystem, err.Error())
			}
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
		DefaultLevel: defaultLevel,
	}
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
