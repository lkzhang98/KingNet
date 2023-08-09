// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package log

import (
	"go.uber.org/zap/zapcore"
)

// Options defines the configuration for the logger.
type Options struct {
	// DisableCaller option to disable the caller information.
	DisableCaller bool
	// DisableStacktrace option to disable the stacktrace information.
	DisableStacktrace bool
	// DisableLevel option to disable the level information.
	DisableLevel bool
	// Level option to set the log level.
	Level string
	// Format option to set the log format.
	// format can be "json" or "console".
	Format string
	// OutputPaths option to set the log output paths.
	// output paths can be a file path or a directory path.
	OutputPaths []string
}

// NewOptions returns a new Options instance with default values.
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
