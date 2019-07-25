// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
//
// Mainly copied from TIDB's log format
package replixlogger

// LogConfig serializes log related config in toml/json.
//
// LogCOnfig contains FileLogConfig where FileLogConfig contains bellowing elements
// 1: Log file name:
//			It should be a valid file location. ReplixLogger will throw error provided
//			file description is a directory
// 2: LogRotate:
//			A boolean value indicates whether to rotate logs
// 3: MaxSize:
//			Max size for a single file, in MB.
// 4: MaxDays:
//			Max log keep days, default is never deleting.
// 5: MaxBackups:
//			Maximum number of old log files to retain.
type LogConfig struct {
	// Log level.
	Level string `toml:"level" json:"level"`
	// Log format. one of json, text, or console.
	Format string `toml:"format" json:"format"`
	// Disable automatic timestamps in output.
	DisableTimestamp bool `toml:"disable-timestamp" json:"disable-timestamp"`
	// File log config.
	File FileLogConfig `toml:"file" json:"file"`
	// SlowQueryFile filename, default to File log config on empty.
	SlowQueryFile string
}
