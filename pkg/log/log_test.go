package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestNewDiscard(t *testing.T) {
	logger := NewDiscard()
	logger.Info("test discard", zap.String("output", "discard"))
}
