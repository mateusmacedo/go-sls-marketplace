package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewZapLogger(t *testing.T) {
	logger, err := NewZapLogger()
	assert.NotNil(t, logger)
	assert.NoError(t, err)
}

func TestZapLogger_Info(t *testing.T) {
	logger := &ZapLogger{
		logger: zap.NewExample().Sugar(),
	}

	logger.Info("This is an info message", "key1", "value1", "key2", "value2")
}

func TestZapLogger_Error(t *testing.T) {
	logger := &ZapLogger{
		logger: zap.NewExample().Sugar(),
	}
	err := errors.New("some error")

	logger.Error("This is an error message", err, "key1", "value1", "key2", "value2")
}
