package utils

import (
	common "QueryLogBuilder4Go/src/common"
	"QueryLogBuilder4Go/src/querylog"
	"math/rand"
)

var alphabetic = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numeric = []rune("0123456789")
var alphanumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var factory *querylog.EventDataFactory = nil
var timeSource common.TimeSource = &common.RealTimeSource{}

func GetCurrentTimeMillis() uint64 {
	return uint64(timeSource.CurrentTimeMillis())
}

func RandAlphaNumeric(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(b)
}

func RandNumeric(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numeric[rand.Intn(len(numeric))]
	}
	return string(b)
}

func RandAlpha(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabetic[rand.Intn(len(alphabetic))]
	}
	return string(b)
}

func GetEventDataFactory() *querylog.EventDataFactory {
	timeSource = &common.RealTimeSource{}
	factory := querylog.NewEventDataFactory(timeSource, "HelloTencent", EMPTY_STRING, nil, nil)
	return factory
}

func GetEventData() querylog.EventData {
	if factory == nil {
		factory = GetEventDataFactory()
	}

	event := factory.CreateEventData()

	event.SetRemoteAddr(querylog.ObtainHostName())
	return event
}

func StartEventData(event querylog.EventData) {
	// Can not happen, but let's still check it
	if timeSource == nil {
		timeSource = &common.RealTimeSource{}
	}

	event.SetStartTime(timeSource.CurrentTimeMillis())
}

func EndEventData(event querylog.EventData) {
	// Can not happen, but let's still check it
	if timeSource == nil {
		timeSource = &common.RealTimeSource{}
	}

	event.SetEndTime(timeSource.CurrentTimeMillis())
	event.RecordProfiledData()
}

func GetThreadSafeEventData() querylog.EventData {
	if factory == nil {
		factory = GetEventDataFactory()
	}

	event := factory.CreateThreadSafeEventData()

	return event
}