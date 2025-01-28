package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitZap initializes the zap logger
func InitZap() {
	var err error

	// Use zap.NewDevelopment() for local development
	// Use zap.NewProduction() for prod development
	Log, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize zap logger: " + err.Error())
	}
	defer Log.Sync() // Flushes buffer, if any
}
