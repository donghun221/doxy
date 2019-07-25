// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
//
// Mainly copied from TIDB's log format
package replixlogger

// FileLogConfig serializes file log related config in toml/json.
type FileLogConfig struct {
	// Log filename, leave empty to disable file log.
	Filename string `toml:"filename" json:"filename"`
	// Is log rotate enabled. TODO.
	LogRotate bool `toml:"log-rotate" json:"log-rotate"`
	// Max size for a single file, in MB.
	MaxSize uint `toml:"max-size" json:"max-size"`
	// Max log keep days, default is never deleting.
	MaxDays uint `toml:"max-days" json:"max-days"`
	// Maximum number of old log files to retain.
	MaxBackups uint `toml:"max-backups" json:"max-backups"`
}


