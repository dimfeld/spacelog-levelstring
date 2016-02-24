package levelstring

import (
	"regexp"
	"strings"

	"github.com/spacemonkeygo/spacelog"
)

// var newLoggerLevels map[spacelog.LogLevel]*regexp.Regexp
var replacer *strings.Replacer

func init() {
	// newLoggerLevels = make(map[spacelog.LogLevel]*regexp.Regexp)
	replacer = strings.NewReplacer(
		".", "\\.",
		"?", "\\?",
		"+", "\\+",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"|", "\\|",
		"*", ".*")
}

// Configure sets all loggers matching a comma-separated glob-style string
// to the requested level.
func Configure(setting string, level spacelog.LogLevel) error {
	if setting == "" {
		return nil
	}

	// Convert from a glob-style string to a regex-style string
	settings := strings.Split(replacer.Replace(setting), ",")
	for i := range settings {
		settings[i] = strings.TrimSpace(settings[i])
	}
	regex, err := regexp.Compile(strings.Join(settings, "|"))
	if err != nil {
		return err
	}
	spacelog.SetLevel(regex, level)

	return nil
}
