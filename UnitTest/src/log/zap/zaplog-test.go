package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	log, _ := zap.NewDevelopment()
	defer log.Sync()

	log.Info("failed to Process",
		zap.String("err", "null exception"),
		zap.Int("attempt", 3),
		zap.Duration("duration", time.Second),
	)

	var logger *zap.Logger
	{
		config := zap.NewDevelopmentConfig()
		//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		config.OutputPaths = []string{
			"./log/zap/zap01.log",
		}

		logger, _ = config.Build()
	}

	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	logger.Info("failed to Process",
		zap.String("err", "null exception"),
		zap.Int("attempt", 3),
		zap.Duration("duration", time.Second),
	)

}
