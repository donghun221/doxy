// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

const (
	MAX_VALUE_SIZE int = 1024
	SCOPE_DELIMITER string = "------------------------------------------------------------------------"
	CRLF string = "\n"
	EOE string = "EOE"
)

// An interface for recording an Aperture/Replix query log entry for a single event.
// A typical event is a GET request, and typical contents would be the total
// time, the time for some sub requests, the fail/succeed outcome, etc.
type EventData interface {
	// Application name which is process name usually
	GetApplicationName() string

	// Please do not set the host name as IP address
	// Use host name as aperture-prod-assembler-0101.cq5.tencent.com
	GetHostName() string

	// Event data name like GET / PUT / DELETE / MIGRATION
	GetOperation() string

	// Event data name like GET / PUT / DELETE / MIGRATION
	SetOperation(string)

	GetStatus() string

	SetStatus(string)

	GetQueryLogStatus() string

	SetQueryLogStatus(string)

	Reset()

	// Timer related
	GetStartTime() int64

	SetStartTime(int64)

	GetEndTime() int64

	SetEndTime(endTime int64)

	// starts a timer. Many timers with the same name can running simultaneously.
	// When more than one timer is running we aggregate time faster than the wall clock.
	StartTimer(string) error

	// Ends a timer event and records a history entry for it.
	EndTimer(string)

	// Adds a sample to a timer aggregator.  Use this method when you
	// have determined the elapsed time yourself
	UpdateTimer(string, int64) error

	UpdateTimerWithSample(string, int64, int64) error

	// Get elapsed time for the specified timer. If no timer
	// exists for the name, null is returned.
	GetTimeElapsed(string) int64

	// Remote Address related
	GetRemoteAddr() string

	SetRemoteAddr(string)

	// Counter related
	// Get the count of the specified counter.
	GetCounter(string) int64

	SetCounter(string, int64)

	// Increments the named counter by the given value. If the counter doesn't
	// exist, then the counter is set to the given value.
	InCCounter(string, int64)

	// Name value pair
	// Adds the provided name and value pair to the event data, replacing the
	// entry for the previous name if it already exists.
	AddNameValuePair(name, value string)

	// appends the value to an existing name value pair.
	// If the value does not exist in the collection of
	// name value pairs, it is added.  If the value already
	// exists, a comma is appended to the existing value
	// and then the new value is appended.
	//
	// If you append more than DEFAULT_MAX_VALUE_COUNT
	// records we will throw out the oldest records.
	AppendNameValuePair(name, value string)

	GetValue(name string) string

	// Finish Event Data
	// records the time of the last event.
	// event time is measured as the time difference
	// between calls to this method.  There are uses
	// where we want to measure the time that we are
	// in mutually exclusive states.  There are situations
	// where determining the next or the previous state is
	// not possible at the same instant.  This method alleviates
	// the problem. You need not know want timer will be starting.
	// You just have to supply the timer that finished.
	//
	// One property of this method is that the sum of all events recorded
	// with this method will equal the sum the time for the entire event data.
	// (As long as start time is not changed)
	FinishCurrentEvent(name string)

	// inserts an event into the event history.
	RecordHistoryEvent(name string)

	RecordProfiledData()

	// Output format related
	ToJsonFormat() string

	ToEventDataFormat() string
}