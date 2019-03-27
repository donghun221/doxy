package querylog

import (
	"bytes"
	"time"
	"strconv"
	"encoding/json"
)

type QueryLogEntryImpl struct {}

func (entry *QueryLogEntryImpl) Format(event *EventDataImpl) string {
	event.addDefaultNameValuePairs()

	builder := bytes.Buffer{}
	builder.WriteString(SCOPE_DELIMITER)
	builder.WriteString(CRLF)

	entry.addPredifinedValues(&builder, event)

	event.appendNameValuePairs(&builder)

	builder.WriteString(EOE)
	return builder.String()
}

func (entry *QueryLogEntryImpl) FormatAsJson(event *EventDataImpl) string {
	val, _ := json.Marshal("Not implemented yet!")
	return string(val)
}

func (entry *QueryLogEntryImpl) addPredifinedValues(builder *bytes.Buffer, event *EventDataImpl) {
	// EndTime
	builder.WriteString("EndTime")
	builder.WriteByte('=')
	builder.WriteString(time.Unix(0, event.GetEndTime()*1000000).String())
	builder.WriteString(CRLF)

	// StartTime
	builder.WriteString("StartTime")
	builder.WriteByte('=')
	builder.WriteString(time.Unix(0, event.GetStartTime()*1000000).String())
	builder.WriteString(CRLF)

	// Time
	builder.WriteString("Time")
	builder.WriteByte('=')
	builder.WriteString(strconv.FormatInt(event.GetEndTime() - event.GetStartTime(), 10)+ " msecs")
	builder.WriteString(CRLF)

	// Hostname
	builder.WriteString("Hostname")
	builder.WriteByte('=')
	builder.WriteString(event.GetHostName())
	builder.WriteString(CRLF)

	// Timing
	builder.WriteString("Timing")
	builder.WriteByte('=')
	event.appendTimers(builder)
	builder.WriteString(CRLF)

	// Counters
	if event.hasCounters() {
		builder.WriteString("Counters")
		builder.WriteByte('=')
		event.appendCounters(builder)
		builder.WriteString(CRLF)
	}

	// Info
	builder.WriteString("Info")
	builder.WriteByte('=')
	builder.WriteString(event.GetOperation())
	builder.WriteString(CRLF)

	// Remote address
	builder.WriteString("RemoteAddress")
	builder.WriteByte('=')
	builder.WriteString(event.GetRemoteAddr())
	builder.WriteString(CRLF)

	// Program
	builder.WriteString("Program")
	builder.WriteByte('=')
	builder.WriteString(event.GetApplicationName())
	builder.WriteString(CRLF)

	// History
	if event.producesHistory() {
		builder.WriteString("History")
		builder.WriteByte('=')
		event.eventHistory.AppendTo(builder)
		builder.WriteString(CRLF)
	}
}