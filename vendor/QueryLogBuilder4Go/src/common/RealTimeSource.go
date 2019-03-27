// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

import (
	"time"
)

// This class implements the TimeSource interface using the time.Now().UnixNano().
//
// Why we need such a method?
// 1: Read current time from system is expensive, so let's wrap time struct for future optimization
// 2: For testing purpose, we will provide a delta function for time pass
type RealTimeSource struct {}

func (source *RealTimeSource) init() {}

// Please optimize this function if it cost system resource
func (source *RealTimeSource) CurrentTimeMillis() int64 {
	return time.Now().UnixNano()/(int64(time.Millisecond))
}