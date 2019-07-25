package querylog

import (
	"time"
	"encoding/binary"
)

func GetSystemTimeMillis() int64 {
	return time.Now().UnixNano()/(int64(time.Millisecond))
}

func Int64ToByteArray(value int64) []byte {
	array := make([]byte, 8)
	binary.LittleEndian.PutUint64(array, uint64(value))

	return array
}