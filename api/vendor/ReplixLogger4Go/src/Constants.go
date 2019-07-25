package replixlogger

import log "github.com/sirupsen/logrus"

const (
	DefaultLogTimeFormat = "2006/01/02 15:04:05.000"
	// DefaultLogMaxSize is the default size of log files.
	DefaultLogMaxSize = 300 // MB
	DefaultLogFormat  = "text"
	DefaultLogLevel   = log.InfoLevel
)
