package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = zap.Must(zap.NewProduction()).Sugar()

func init() {
	stdout := zapcore.AddSync(os.Stdout)

	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	switch os.Getenv("LOGGER_LEVEL") {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warning", "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	jsonEncoder := zapcore.NewJSONEncoder(productionCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	devLogger := zap.New(zapcore.NewCore(consoleEncoder, stdout, level))
	prodLogger := zap.New(zapcore.NewCore(jsonEncoder, stdout, level))

	// TODO: Fix this
	if os.Getenv("APP_ENV") != "production" {
		logger = devLogger.Sugar()
	} else {
		logger = prodLogger.Sugar()
	}

}

func GetLogger() *zap.SugaredLogger {
	return logger
}
