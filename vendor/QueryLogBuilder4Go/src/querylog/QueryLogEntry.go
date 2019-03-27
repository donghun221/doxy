// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

// Interface for immutable query log entries.
type QueryLogEntry interface {
	Format(*EventDataImpl) string
}
