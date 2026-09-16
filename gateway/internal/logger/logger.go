// internal/logger/logger.go
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/gateway.log",
		MaxSize:    50, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	})
	consoleWriter := zapcore.AddSync(os.Stdout)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), fileWriter, zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), consoleWriter, zapcore.DebugLevel),
	)

	zap.ReplaceGlobals(zap.New(core))
}

func Sync() {
	_ = zap.L().Sync()
}
