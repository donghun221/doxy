// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

/**
 * Interface for getting the current time.
 * This is handy for testing.
 * Also, the implementation that uses the real system
 * clock can guarantee that time does not go backwards
 * as it sometimes does at amazon.
 *
 */
type TimeSource interface {
	// get the current time millis.
	CurrentTimeMillis() int64
}
