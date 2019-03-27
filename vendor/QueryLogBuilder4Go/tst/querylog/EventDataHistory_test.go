package tst

import (
	"testing"
	"QueryLogBuilder4Go/src/querylog"
	"QueryLogBuilder4Go/tst/base"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestEventDataHistory_WithBigHistory(t *testing.T) {
	history := querylog.NewEventHistory()

	str := tst.RandString(querylog.MAX_HISTORY_LENGTH + 1)
	history.Elapsed(str, tst.ONE)
	assert.Equal(t, querylog.TRUNCATED_STRING, history.ToString())
}

func TestEventDataHistory_WithOneEvent(t *testing.T) {
	history := querylog.NewEventHistory()

	history.Elapsed(tst.MOCK_EVENT_NAME, tst.ONE)
	assert.Equal(t, tst.MOCK_EVENT_NAME + ":1", history.ToString())
}

func TestEventDataHistory_WithTwoEvent(t *testing.T) {
	history := querylog.NewEventHistory()

	history.Elapsed("eventA", tst.ONE)
	history.Elapsed("eventB", tst.THREE)

	assert.Equal(t, "eventA:1,eventB:2", history.ToString())
}

func TestEventDataHistory_WithThreeEvent_Overflow_One(t *testing.T) {
	history := querylog.NewEventHistory()

	history.Elapsed("eventA", tst.ONE)
	history.Elapsed("eventB", tst.THREE)
	history.Elapsed(tst.RandString(querylog.MAX_HISTORY_LENGTH), tst.THREE)
	assert.Equal(t, "eventA:1,eventB:2,TRUNCATED", history.ToString())
}

func TestEventDataHistory_WithThreeEvent_Overflow_Two(t *testing.T) {
	history := querylog.NewEventHistory()

	history.Elapsed("eventA", tst.ONE)
	history.Elapsed("eventB", tst.THREE)
	history.Elapsed(tst.RandString(querylog.MAX_HISTORY_LENGTH), tst.THREE)
	history.Elapsed("eventC", tst.THREE)
	assert.Equal(t, "eventA:1,eventB:2,TRUNCATED", history.ToString())
}

func TestEventDataHistory_AppendTo(t *testing.T) {
	history := querylog.NewEventHistory()

	history.Elapsed("eventA", tst.ONE)
	history.Elapsed("eventB", tst.THREE)
	history.Elapsed("eventC", tst.THREE)

	buffer := bytes.Buffer{}
	history.AppendTo(&buffer)

	assert.Equal(t, "eventA:1,eventB:2,eventC:0", buffer.String())
}