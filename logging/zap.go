package logging

import "go.uber.org/zap"

func NewZapLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()

	return logger, err
}
