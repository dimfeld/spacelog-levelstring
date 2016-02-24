package levelstring

import (
	"regexp"
	"strings"

	"github.com/spacemonkeygo/spacelog"
)

// Configure sets all loggers matching a comma-separated glob-style string
// to the requested level.
func Configure(setting string, level spacelog.LogLevel) error {
	if setting == "" {
		return nil
	}

	settings := strings.Split(setting, ",")
	for _, setting := range settings {
		// Convert from a glob-style string to a regex-style string
		setting = strings.Replace(strings.TrimSpace(setting), "*", ".*", -1)
		regex, err := regexp.Compile(setting)
		if err != nil {
			return err
		}
		spacelog.SetLevel(regex, level)
	}

	return nil
}
