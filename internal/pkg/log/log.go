// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper around zap.Logger that implements the Logger interface.
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Panic(msg ...interface{})
	Fatal(msg ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

// KLogger is a wrapper around zap.SugaredLogger that implements the Logger interface.
type KLogger struct {
	z *zap.Logger
}

// Ensure the Klogger implements Logger interface.
var _ Logger = &KLogger{}

var (
	// mu is used to synchronize access to the global logger.
	mu sync.Mutex

	// std Logger global variables.
	std = NewLogger(NewOptions())
)

// Init initializes the global logger with the given options.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// NewLogger returns a new KLogger with the given options.
func NewLogger(opts *Options) *KLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// create the default logger level
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// if the log level is not valid, use the default level
		zapLevel = zapcore.InfoLevel
	}

	// create a logger
	encoderConfig := zap.NewProductionEncoderConfig()
	// defines the message
	encoderConfig.MessageKey = "message"
	// defines the timestamp
	encoderConfig.TimeKey = "timestamp"
	// defines the timestamp serialized function
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// set the timestamp serialized function
	// set the timestamp interval
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	// use cfg to create a new zap.Logger
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &KLogger{z: z}

	// redirect the standard library log.Logger's info level output to zap.Logger
	zap.RedirectStdLog(z)

	return logger
}

// Sync calls the underlying zap.Logger's Sync method, which flushes the cached logs to disk.
func Sync() { std.Sync() }

func (k *KLogger) Sync() {
	_ = k.z.Sync()
}

func Debug(msg ...interface{}) {
	std.z.Sugar().Debug(msg)
}

func (k *KLogger) Debug(msg ...interface{}) {
	k.z.Sugar().Debug(msg)
}

func (k *KLogger) Debugf(format string, v ...interface{}) {
	k.z.Sugar().Debugf(format, v)
}

func Debugf(format string, v ...interface{}) {
	std.z.Sugar().Debugf(format, v)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (k *KLogger) Debugw(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Debugw(msg, keysAndValues...)
}

func Info(msg ...interface{}) {
	std.z.Sugar().Info(msg)
}

func (k *KLogger) Info(msg ...interface{}) {
	k.z.Sugar().Info(msg)
}

func (k *KLogger) Infof(format string, v ...interface{}) {
	k.z.Sugar().Infof(format, v)
}

func Infof(format string, v ...interface{}) {
	std.z.Sugar().Infof(format, v)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (k *KLogger) Infow(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Infow(msg, keysAndValues...)
}

func Warn(msg ...interface{}) {
	std.z.Sugar().Warn(msg)
}

func (k *KLogger) Warn(msg ...interface{}) {
	k.z.Sugar().Warn(msg)
}

func (k *KLogger) Warnf(format string, v ...interface{}) {
	k.z.Sugar().Warnf(format, v)
}

func Warnf(format string, v ...interface{}) {
	std.z.Sugar().Warnf(format, v)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (k *KLogger) Warnw(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Warnw(msg, keysAndValues...)
}

func Error(msg ...interface{}) {
	std.z.Sugar().Error(msg)
}

func (k *KLogger) Error(msg ...interface{}) {
	k.z.Sugar().Error(msg)
}

func (k *KLogger) Errorf(format string, v ...interface{}) {
	k.z.Sugar().Errorf(format, v)
}

func Errorf(format string, v ...interface{}) {
	std.z.Sugar().Errorf(format, v)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (k *KLogger) Errorw(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Errorw(msg, keysAndValues...)
}

func Panic(msg ...interface{}) {
	std.z.Sugar().Panic(msg)
}

func (k *KLogger) Panic(msg ...interface{}) {
	k.z.Sugar().Panic(msg)
}

func (k *KLogger) Panicf(format string, v ...interface{}) {
	k.z.Sugar().Panicf(format, v)
}

func Panicf(format string, v ...interface{}) {
	std.z.Sugar().Panicf(format, v)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (k *KLogger) Panicw(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Panicw(msg, keysAndValues...)
}

func Fatal(msg ...interface{}) {
	std.z.Sugar().Fatal(msg)
}

func (k *KLogger) Fatal(msg ...interface{}) {
	k.z.Sugar().Fatal(msg)
}

func (k *KLogger) Fatalf(format string, v ...interface{}) {
	k.z.Sugar().Fatalf(format, v)
}

func Fatalf(format string, v ...interface{}) {
	std.z.Sugar().Fatalf(format, v)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (k *KLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	k.z.Sugar().Fatalw(msg, keysAndValues...)
}

// clone deep copies the KLogger.
func (k *KLogger) clone() *KLogger {
	lc := *k
	return &lc
}
