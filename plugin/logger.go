package plugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/ipfs/go-ipfs/plugin"
	logging "github.com/ipfs/go-log"
)

// LoggerPlugin controls which ipfs subsystems loggers are enabled
// These subsystems will be defined in env variable:
// eg: LOGGER_SUBSYSTEMS_WHITELIST = "dht,relay,corerepo"
type LoggerPlugin struct{}

var _ plugin.Plugin = &LoggerPlugin{}

const loggerSubsystemsEnv = "LOGGER_SUBSYSTEMS_WHITELIST"

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
	whitelistedSubsystems := os.Getenv(loggerSubsystemsEnv)
	// If no subsystems given, exit with default settings
	if whitelistedSubsystems == "" {
		return nil
	}

	logging.SetAllLoggers(logging.LevelFatal)
	for _, s := range strings.Split(whitelistedSubsystems, ",") {
		subsystem := strings.TrimSpace(s)
		err := logging.SetLogLevel(subsystem, "info")
		if err != nil {
			fmt.Printf("[Warning] Set log level failed for subsystem: %s. Error: %s", subsystem, err.Error())
		}
	}

	return nil
}
