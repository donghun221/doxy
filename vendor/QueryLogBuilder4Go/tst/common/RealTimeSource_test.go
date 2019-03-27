package common

import (
	"testing"
	"QueryLogBuilder4Go/src/common"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestRealTimeSource_HappyCase(t *testing.T) {
	timeSource := querylog.RealTimeSource{}

	currentMilli := time.Now().UnixNano()/int64(time.Millisecond)

	timeSourceMilli := timeSource.CurrentTimeMillis()

	assert.True(t, currentMilli <= timeSourceMilli)
}