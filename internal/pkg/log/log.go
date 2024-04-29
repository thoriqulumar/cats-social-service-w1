package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new zap logger with the given log level, service name and environment.
func New(level zapcore.Level, serviceID, version string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	config.DisableStacktrace = true
	config.Sampling = nil
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.InitialFields = map[string]interface{}{
		"service": serviceID,
		"version": version,
	}

	zapLog, err := config.Build()
	if err != nil {
		return nil, err
	}

	return zapLog, nil
}
