// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

import (
	"bytes"
	"strconv"
)

// This class builds a string that records the history of the event.
// To record an event just call elapsed with the name of the event, and
// the currentTime.  Note you have to send in the time, because we
// know that the EventData class has just gotten the current time,
// this is an optimization to reduce the number of calls to get the time.
//
// The resulting string grows each time the elapsed() method is called;
// "action:deltatime" is appended to it, where "action" is the action parameter
// to elapsed, and deltatime is an integer indicating the difference between the
// time parameter to elapsed, and the last time parameter to elapsed. (The first
// item in the string has a deltatime of 0.)
//
// If a call to elapsed() would grow the string to be larger than the maxlength,
// "TRUNCATED" is appended and further calls to elapsed() are ignored.
//
// A resulting string might look like:
//      e-resolve:0,e-connect:35,e-request:11,e-close:9
// or
//      e-wake:0,e-clean:2,e-eat:51,e-walk:9,e-read:11,TRUNCATED
//
// Note that "e-" and "s-" are used to indicate the end and start of an event,
// and are passed in as the event name by the primary consumer of this class,
// EventDataImpl.
const (
	TRUNCATED_STRING string = "TRUNCATED"
	COMMA_TRUNCATED_STRING string = ",TRUNCATED"
	MAX_HISTORY_LENGTH int = 1024
)

type EventHistory struct {
	// representing our accumulated history string.
	builder 		bytes.Buffer
	// Set true if we've ever truncated; avoids further computation on lengths.
	truncated 		bool
    // "time" parameter from the previous call to elapsed() method; used to compute
	// deltas to the next time it's called. Those deltas are stored in the string as
	// we build it.
	previousMillis int64
}

func NewEventHistory() *EventHistory {
	his := EventHistory{}

	his.builder = bytes.Buffer{}
	his.truncated = false
	his.previousMillis = 0
	return &his
}

// Appends an elapsed time entry to the history.
// if the history record would exceed the maxHistorySize,
// then the history, is marked as truncated.
//
// action: the string name of the action for which history is being recorded
// time: the time of the history event. This implementation is unit agnostic, but most callers use milliseconds.
func (his *EventHistory) Elapsed(action string, time int64) {
	// the history is already truncated
	if his.truncated {
		return
	}

	// see if we have room for more history.
	// NOTE, we always have to leave enough space to append the TRUNCATED_STRING
	// because the next call may need to append it. (we never know if this is
	// the last call to elapsed)
	length := his.builder.Len()
	elapsed := strconv.FormatInt(time - his.previousMillis, 10)

	size := len(action) + 1 + len(elapsed)
	if length > 0 {
		size++
	}

	if length + size + 1 + len(TRUNCATED_STRING) > MAX_HISTORY_LENGTH {
		his.truncated = true
		if length > 0 {
			// we have something in the string and adding more would've
			// put us over our limit, so just mark the string truncated
			his.builder.WriteString(COMMA_TRUNCATED_STRING)
		} else {
			// we have nothing in the string and we were asked to add
			// something so large that we'd immediately be over the limit;
			// we'll immediately go TRUNCATED in this case.
			his.builder.WriteString(TRUNCATED_STRING)
		}
		return
	}

	// save the previous time
	his.previousMillis = time

	// append a comma if we've got a previous string
	if length > 1 {
		his.builder.WriteByte(',')
	}

	// append our next action, a colon, and its delta time
	his.builder.WriteString(action)
	his.builder.WriteByte(':')
	his.builder.WriteString(elapsed)
}

// sets up this object for reuse
func (his *EventHistory) Clear() {
	his.builder.Reset()
	his.previousMillis = 0
	his.truncated = false
}

// appends the event history data to the given byte buffer
func (his *EventHistory) AppendTo(buffer *bytes.Buffer) {
	buffer.Write(his.builder.Bytes())
}

// returns the current history string
func (his *EventHistory) ToString() string {
	return his.builder.String()
}

// the number of characters in the byte buffer
func (his *EventHistory) GetHistoryLength() int {
	return his.builder.Len()
}
