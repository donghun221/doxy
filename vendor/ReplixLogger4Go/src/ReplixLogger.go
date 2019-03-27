// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
//
// Mainly copied from TIDB's log format
package replixlogger

import (
	"github.com/juju/errors"
	"strings"

	log "github.com/sirupsen/logrus"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
)

// SlowQueryLogger is used to log slow query, InitLogger will modify it according to config file.
var SlowQueryLogger = log.StandardLogger()

// InitLogger initializes logger.
func InitLogger(cfg *LogConfig) error {
	log.SetLevel(stringToLogLevel(cfg.Level))
	log.AddHook(&ContextHook{})

	if cfg.Format == "" {
		cfg.Format = DefaultLogFormat
	}
	formatter := stringToLogFormatter(cfg.Format, cfg.DisableTimestamp)
	log.SetFormatter(formatter)

	if len(cfg.File.Filename) != 0 {
		if err := InitFileLog(&cfg.File, nil); err != nil {
			return errors.Trace(err)
		}
	}

	if len(cfg.SlowQueryFile) != 0 {
		SlowQueryLogger = log.New()
		tmp := cfg.File
		tmp.Filename = cfg.SlowQueryFile
		if err := InitFileLog(&tmp, SlowQueryLogger); err != nil {
			return errors.Trace(err)
		}
		hooks := make(log.LevelHooks)
		hooks.Add(&ContextHook{})
		SlowQueryLogger.Hooks = hooks
		slowQueryFormatter := stringToLogFormatter(cfg.Format, cfg.DisableTimestamp)
		ft, ok := slowQueryFormatter.(*textFormatter)
		if ok {
			ft.EnableEntryOrder = true
		}
		SlowQueryLogger.Formatter = slowQueryFormatter
	}

	return nil
}

func stringToLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "fatal":
		return log.FatalLevel
	case "error":
		return log.ErrorLevel
	case "warn", "warning":
		return log.WarnLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	}
	return DefaultLogLevel
}

func stringToLogFormatter(format string, disableTimestamp bool) log.Formatter {
	switch strings.ToLower(format) {
	case "text":
		return &textFormatter{
			DisableTimestamp: disableTimestamp,
		}
	case "json":
		return &log.JSONFormatter{
			TimestampFormat:  DefaultLogTimeFormat,
			DisableTimestamp: disableTimestamp,
		}
	case "console":
		return &log.TextFormatter{
			FullTimestamp:    true,
			TimestampFormat:  DefaultLogTimeFormat,
			DisableTimestamp: disableTimestamp,
		}
	case "highlight":
		return &textFormatter{
			DisableTimestamp: disableTimestamp,
			EnableColors:     true,
		}
	default:
		return &textFormatter{}
	}
}

// initFileLog initializes file based logging options.
func InitFileLog(cfg *FileLogConfig, logger *log.Logger) error {
	if st, err := os.Stat(cfg.Filename); err == nil {
		if st.IsDir() {
			return errors.New("can't use directory as log file name")
		}
	}
	if cfg.MaxSize == 0 {
		cfg.MaxSize = DefaultLogMaxSize
	}

	// use lumberjack to logrotate
	output := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    int(cfg.MaxSize),
		MaxBackups: int(cfg.MaxBackups),
		MaxAge:     int(cfg.MaxDays),
		LocalTime:  true,
	}

	// Move to logrus default logger if logger is nil
	if logger == nil {
		log.SetOutput(output)
	} else {
		logger.Out = output
	}
	return nil
}
