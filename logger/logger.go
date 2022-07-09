package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// access outside the function, l (private) L (public)
var log *zap.Logger

// initialize zap logger
// to use an imported package it needs to be initialized first. it's done by Golang's runtime system
// and consists of (order matters):
// initialization of imported packages (recursive definition)
// computing and assigning initial value for variables declared in a package block.
// EXECUTING INIT FUNCTIONS INSIDE THE PACKAGE
func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	// looked into the code and did it manually, to create our customer logger
	log, err = config.Build(zap.AddCallerSkip(1))

	// log, err = zap.NewProduction(zap.AddCallerSkip(1)) // wrapping it only one level
	if err != nil {
		panic(err)
	}
}

// private to isolate dependency to our code not the package code
// caller from looger.go not from main.go, want to change so that
// it tells us from where, main.go is where originally
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// want to always wrap to make it depended to our code not the zap in case
// we ever want to change zap, the cod ein main.go won't break and have to
// be changed.
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
