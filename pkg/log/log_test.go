package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, logger.level, testCase.level)
		})
	}
}
