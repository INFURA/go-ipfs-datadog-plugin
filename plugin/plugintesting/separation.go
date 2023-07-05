package plugintesting

import (
	"os"
	"testing"
)

const (
	envVarIntegrationTests = "INTEGRATION"
	integrationTestsMsg    = "set the " + envVarIntegrationTests + " environment variable to run integration tests (see README.md for more details)"
)

func Integration(t *testing.T) {
	t.Helper()

	if _, ok := os.LookupEnv(envVarIntegrationTests); !ok {
		t.Skip(integrationTestsMsg)
	}
}
