package log

import (
	"testing"
)

func TestSetLoggerLevels(t *testing.T) {
	table := []struct {
		level  LoggerLevel
		expect string
	}{
		{Info, "INFO"},
		{Debug, "DEBUG"},
		{Warning, "WARNING"},
		{Error, "ERROR"},
	}

	for _, testCase := range table {
		t.Run(testCase.expect, func(t *testing.T) {
			logger := NewDefaultLogger(testCase.level)
			if logger.level != testCase.level {
				t.Errorf("expected %v, got %v", testCase.expect, logger.level)
			}
		})
	}
}
