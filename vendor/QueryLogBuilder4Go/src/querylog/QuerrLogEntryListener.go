// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

// Interface for objects interested in query log entries
type QueryLogEntryListener interface {
	// Notifies the listener of query log entry
	notify(QueryLogEntry)
}