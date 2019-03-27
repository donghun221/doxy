package tst

import (
	"testing"
	"QueryLogBuilder4Go/src/querylog"
	"github.com/stretchr/testify/assert"
	"math/rand"
	common "QueryLogBuilder4Go/src/common"
)

const (
	ZERO int64 = 0
	ONE int64 = 1
	TWO int64 = 2
	THREE int64 = 3
	FOUR int64 = 4
	FIVE int64 = 5
	SIX int64 = 6
	SEVEN int64 = 7
	EIGHT int64 = 8
	NINE int64 = 9
	TEN int64 = 10
	ELEVEN int64 = 11

	MOCK_EVENT_NAME string = "mock_event"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GetTimerAggregator(t *testing.T) *querylog.TimerAggregator {
	agg, err := querylog.NewTimerAggregator(MOCK_EVENT_NAME)

	assert.Nil(t, err)
	assert.NotNil(t, agg)
	return agg
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type TestQueryLogEntryListener struct {
	entry querylog.QueryLogEntry
}

func (listener *TestQueryLogEntryListener) notify(entry querylog.QueryLogEntry) {
	listener.entry = entry
}

func GetEventDataFactoryDefault() *querylog.EventDataFactory {
	factory := querylog.NewEventDataFactoryDefault()

	return factory
}

func GetEventDataFactoryWithTimeSource(timeSource common.TimeSource) *querylog.EventDataFactory {
	factory := querylog.NewEventDataFactory(timeSource, common.EMPTY_STRING, common.EMPTY_STRING, nil, nil)

	return factory
}

func GetNoopEventData() querylog.EventData {
	return querylog.NoopEventData{}
}