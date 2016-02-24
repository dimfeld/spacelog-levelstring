package levelstring

import (
	"fmt"
	"testing"

	"github.com/spacemonkeygo/spacelog"
)

func TestConfigure(t *testing.T) {
	config := spacelog.SetupConfig{
		Output: "stderr",
		Level:  "warn",
	}
	spacelog.MustSetup("spacelog-debug_test", config)

	logger1Name := "aaa"
	logger1 := spacelog.GetLoggerNamed(logger1Name)
	logger2Name := "bbb1"
	logger2 := spacelog.GetLoggerNamed(logger2Name)
	logger3Name := "bbb2"
	logger3 := spacelog.GetLoggerNamed(logger3Name)

	check := func(l *spacelog.Logger, name string, expectEnabled bool) {
		enabled := l.DebugEnabled()
		if enabled != expectEnabled {
			t.Errorf("%s DebugEnabled expected %v, saw %v", name, expectEnabled, enabled)
		}
	}

	t.Log("Initial setup")
	check(logger1, logger1Name, false)
	check(logger2, logger2Name, false)
	check(logger3, logger3Name, false)

	t.Log("Enable b*")
	Configure("b*", spacelog.Debug)

	check(logger1, logger1Name, false)
	check(logger2, logger2Name, true)
	check(logger3, logger3Name, true)

	t.Log("Disable bbb2")
	Configure("bbb2", spacelog.Info)
	check(logger1, logger1Name, false)
	check(logger2, logger2Name, true)
	check(logger3, logger3Name, false)

	t.Log("Enable *")
	Configure("*", spacelog.Debug)
	check(logger1, logger1Name, true)
	check(logger2, logger2Name, true)
	check(logger3, logger3Name, true)

	t.Log("Disable bbb1, bbb2")
	Configure("bbb1, bbb2", spacelog.Info)
	check(logger1, logger1Name, true)
	check(logger2, logger2Name, false)
	check(logger3, logger3Name, false)

	t.Log("Disable aaa,bbb1,bbb2")
	Configure("bbb1,aaa,bbb2", spacelog.Info)
	check(logger1, logger1Name, false)
	check(logger2, logger2Name, false)
	check(logger3, logger3Name, false)

	t.Log("Enable a*, b*")
	Configure("a*, b*", spacelog.Debug)
	check(logger1, logger1Name, true)
	check(logger2, logger2Name, true)
	check(logger3, logger3Name, true)

	t.Log("Disable *b*")
	Configure("*b1*", spacelog.Info)
	check(logger1, logger1Name, true)
	check(logger2, logger2Name, false)
	check(logger3, logger3Name, true)
}

func ExampleConfigure() {
	spacelog.MustSetup("configure-example", spacelog.SetupConfig{
		Output: "stdout",
		Level:  "info",
		Format: "{{.Level}} {{.LoggerName}} - {{.Message}}",
	})

	inLogger := spacelog.GetLoggerNamed("input")
	inputVerboseLogger := spacelog.GetLoggerNamed("input:verbose")
	coreLogger := spacelog.GetLoggerNamed("core")
	outLogger := spacelog.GetLoggerNamed("output")
	outputVerboseLogger := spacelog.GetLoggerNamed("output:verbose")

	print := func() {
		// Only the enabled facilities will actually print
		inLogger.Debug("input logger debug enabled")
		inputVerboseLogger.Debug("input:verbose logger debug enabled")
		coreLogger.Debug("core logger debug enabled")
		outLogger.Debug("output logger debug enabled")
		outputVerboseLogger.Debug("output:verbose logger debug enabled")
		fmt.Println()
	}

	coreLogger.Info("Enabling input* DEBUG")
	Configure("input*", spacelog.Debug)
	print()

	coreLogger.Info("Enabling *verbose DEBUG")
	Configure("*verbose", spacelog.Debug)
	print()

	coreLogger.Info("Enabling all DEBUG")
	Configure("*", spacelog.Debug)
	print()

	coreLogger.Info("Setting input:verb*, output:verbose back to WARN")
	Configure("input:verb*, output:verbose", spacelog.Warning)
	print()

	// Output:
	// INFO core - Enabling input* DEBUG
	// DEBUG input - input logger debug enabled
	// DEBUG input:verbose - input:verbose logger debug enabled
	//
	// INFO core - Enabling *verbose DEBUG
	// DEBUG input - input logger debug enabled
	// DEBUG input:verbose - input:verbose logger debug enabled
	// DEBUG output:verbose - output:verbose logger debug enabled
	//
	// INFO core - Enabling all DEBUG
	// DEBUG input - input logger debug enabled
	// DEBUG input:verbose - input:verbose logger debug enabled
	// DEBUG core - core logger debug enabled
	// DEBUG output - output logger debug enabled
	// DEBUG output:verbose - output:verbose logger debug enabled
	//
	// INFO core - Setting input:verb*, output:verbose back to WARN
	// DEBUG input - input logger debug enabled
	// DEBUG core - core logger debug enabled
	// DEBUG output - output logger debug enabled
}
