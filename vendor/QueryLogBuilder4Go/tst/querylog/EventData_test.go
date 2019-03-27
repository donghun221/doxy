package tst

import (
	"testing"
	"time"
	"QueryLogBuilder4Go/src/querylog"
	"QueryLogBuilder4Go/tst/base"
	"github.com/stretchr/testify/assert"
	"unsafe"
	"strings"
)

func TestEventData_WithNothingToBeSet(t *testing.T) {
	expected := GetNothingSetStr()

	event := tst.GetEventDataFactoryDefault().CreateEventData()
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_WithNothingToBeSet_AndReset(t *testing.T) {
	expected := GetNothingSetStr()

	event := tst.GetEventDataFactoryDefault().CreateEventData()
	assert.Equal(t, expected, event.ToEventDataFormat())

	event.Reset()
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_WithNoGrowth(t *testing.T) {
	expected := GetNothingSetStr()

	event := tst.GetEventDataFactoryDefault().CreateEventData()
	assert.Equal(t, expected, event.ToEventDataFormat())

	event.Reset()
	assert.Equal(t, expected, event.ToEventDataFormat())
	firstSize := unsafe.Sizeof(event)

	event.Reset()
	assert.Equal(t, expected, event.ToEventDataFormat())
	secondSize := unsafe.Sizeof(event)

	assert.Equal(t, firstSize, secondSize)
}

func TestEventData_SetInstanceDefault(t *testing.T) {
	expected := querylog.SCOPE_DELIMITER + querylog.CRLF +
		"EndTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"StartTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"Time=0 msecs" + querylog.CRLF +
		"Hostname=" + querylog.ObtainHostName() + querylog.CRLF +
		"Timing=" + querylog.CRLF +
		"Info=" + querylog.CRLF +
		"RemoteAddress=" + querylog.CRLF +
		"Program=" + querylog.UNKNOWN_APPLICATION + querylog.CRLF +
		"History=" + querylog.CRLF +
		"Instance=someInstance" + querylog.CRLF +
		"EOE"

	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.AddNameValuePair("Instance", "someInstance")
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_SetStatus(t *testing.T) {
	expected := querylog.SCOPE_DELIMITER + querylog.CRLF +
		"EndTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"StartTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"Time=0 msecs" + querylog.CRLF +
		"Hostname=" + querylog.ObtainHostName() + querylog.CRLF +
		"Timing=" + querylog.CRLF +
		"Info=" + querylog.CRLF +
		"RemoteAddress=" + querylog.CRLF +
		"Program=" + querylog.UNKNOWN_APPLICATION + querylog.CRLF +
		"History=" + querylog.CRLF +
		"Status=testStatus" + querylog.CRLF +
		"EOE"

	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.SetStatus("testStatus")
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_SetOperation(t *testing.T) {
	expected := querylog.SCOPE_DELIMITER + querylog.CRLF +
		"EndTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"StartTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"Time=0 msecs" + querylog.CRLF +
		"Hostname=" + querylog.ObtainHostName() + querylog.CRLF +
		"Timing=" + querylog.CRLF +
		"Info=testOperation" + querylog.CRLF +
		"RemoteAddress=" + querylog.CRLF +
		"Program=" + querylog.UNKNOWN_APPLICATION + querylog.CRLF +
		"History=" + querylog.CRLF +
		"Operation=testOperation" + querylog.CRLF +
		"EOE"

	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.SetOperation("testOperation")
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_Reset(t *testing.T) {
	expected := GetNothingSetStr()

	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.SetOperation("testOperation")
	event.SetStatus("testStatus")

	assert.NotEqual(t, expected, event.ToEventDataFormat())

	event.Reset()
	assert.Equal(t, expected, event.ToEventDataFormat())
}

func TestEventData_SetCounter(t *testing.T) {
	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.SetCounter("counter1", tst.ONE)
	event.SetCounter("counter2", tst.TWO)
	event.SetCounter("counter3", tst.THREE)

	format := event.ToEventDataFormat()

	assert.True(t, strings.Contains(format, "counter1=1"))
	assert.True(t, strings.Contains(format, "counter2=2"))
	assert.True(t, strings.Contains(format, "counter3=3"))
}

func TestEventData_NameValuePairs(t *testing.T) {
	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.AddNameValuePair("Name1", "Value1")
	event.AddNameValuePair("Name2", "Value2")
	event.AddNameValuePair("Name3", "Value3")

	format := event.ToEventDataFormat()

	assert.True(t, strings.Contains(format, "Name1=Value1"))
	assert.True(t, strings.Contains(format, "Name2=Value2"))
	assert.True(t, strings.Contains(format, "Name3=Value3"))
}

func TestEventData_FinishEvent_One(t *testing.T) {
	timeSource := tst.TestTimeSource{}

	timeSource.TimeToReturn = 110
	event := tst.GetEventDataFactoryWithTimeSource(&timeSource).CreateEventData()
	event.SetStartTime(110)
	event.FinishCurrentEvent("event1")

	timeSource.TimeToReturn = 119
	event.SetEndTime(119)

	str := event.ToEventDataFormat()
	assert.True(t, strings.Contains(str, "event1:0/1"))
	assert.True(t, strings.Contains(str, "eventSlop:9/1"))
	assert.True(t, strings.Contains(str, "start:110"))
	assert.True(t, strings.Contains(str, "event1:0"))
	assert.True(t, strings.Contains(str, "eventSlop:9"))
	assert.True(t, strings.Contains(str, "end:0"))
}

func TestEventData_FinishEvent_Two(t *testing.T) {
	timeSource := tst.TestTimeSource{}

	event := tst.GetEventDataFactoryWithTimeSource(&timeSource).CreateEventData()
	event.SetStartTime(100)

	timeSource.TimeToReturn = 110
	event.FinishCurrentEvent("event1")

	timeSource.TimeToReturn = 119
	event.FinishCurrentEvent("event2")

	timeSource.TimeToReturn = 127
	event.SetEndTime(127)

	str := event.ToEventDataFormat()
	assert.True(t, strings.Contains(str, "event1:10/1"))
	assert.True(t, strings.Contains(str, "event2:9/1"))
	assert.True(t, strings.Contains(str, "eventSlop:8/1"))

}

func TestEventData_FinishEvent_Multiple(t *testing.T) {
	timeSource := tst.TestTimeSource{}

	event := tst.GetEventDataFactoryWithTimeSource(&timeSource).CreateEventData()
	event.SetStartTime(100)

	timeSource.TimeToReturn = 110
	event.FinishCurrentEvent("event1")

	timeSource.TimeToReturn = 119
	event.FinishCurrentEvent("event2")

	timeSource.TimeToReturn = 127
	event.FinishCurrentEvent("event1")

	timeSource.TimeToReturn = 137
	event.SetEndTime(137)

	str := event.ToEventDataFormat()
	assert.True(t, strings.Contains(str, "event1:18/2"))
	assert.True(t, strings.Contains(str, "event2:9/1"))
	assert.True(t, strings.Contains(str, "eventSlop:10/1"))

	assert.True(t, strings.Contains(str, "start:100"))
	assert.True(t, strings.Contains(str, "event1:10"))
	assert.True(t, strings.Contains(str, "event2:9"))
	assert.True(t, strings.Contains(str, "event1:8"))
	assert.True(t, strings.Contains(str, "eventSlop:10"))
	assert.True(t, strings.Contains(str, "end:0"))
}

func TestEventData_FinishEvent_Multiple_NoSlop(t *testing.T) {
	timeSource := tst.TestTimeSource{}

	event := tst.GetEventDataFactoryWithTimeSource(&timeSource).CreateEventData()
	event.SetStartTime(100)

	timeSource.TimeToReturn = 110
	event.FinishCurrentEvent("event1")

	timeSource.TimeToReturn = 119
	event.FinishCurrentEvent("event2")

	timeSource.TimeToReturn = 127
	event.FinishCurrentEvent("event1")
	event.SetEndTime(127)

	str := event.ToEventDataFormat()
	assert.True(t, strings.Contains(str, "event1:18/2"))
	assert.True(t, strings.Contains(str, "event2:9/1"))
}

func TestEventData_FinishEvent_AppendNameValuePair(t *testing.T) {
	timeSource := tst.TestTimeSource{}

	event := tst.GetEventDataFactoryWithTimeSource(&timeSource).CreateEventData()
	event.SetStartTime(100)
	timeSource.TimeToReturn = 127

	event.AddNameValuePair("name", "value")
	event.AppendNameValuePair("name", "value2")
	event.AddNameValuePair("name2", "XXvalue")
	event.AddNameValuePair("name2", "XXvalue2")

	event.SetEndTime(127)

	str := event.ToEventDataFormat()

	assert.True(t, strings.Contains(str, "start:100"))
	assert.True(t, strings.Contains(str, "end:27"))
}

func TestEventData_OpenTimer(t *testing.T) {
	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.StartTimer("ticky")

	// do not close the timer, but finish event data
	event.RecordProfiledData()

	assert.True(t, strings.Contains(event.ToEventDataFormat(), "Timing=ticky:"))
}

func TestEventData_UpdateTimer(t *testing.T) {
	event := tst.GetEventDataFactoryDefault().CreateEventData()

	event.UpdateTimer("ticky", 250)

	// do not close the timer, but finish event data
	event.RecordProfiledData()

	assert.True(t, strings.Contains(event.ToEventDataFormat(), "Timing=ticky:250/1"))
}

func GetNothingSetStr() string {
	return querylog.SCOPE_DELIMITER + querylog.CRLF +
		"EndTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"StartTime=" + time.Unix(0,0).String() + querylog.CRLF +
		"Time=0 msecs" + querylog.CRLF +
		"Hostname=" + querylog.ObtainHostName() + querylog.CRLF +
		"Timing=" + querylog.CRLF +
		"Info=" + querylog.CRLF +
		"RemoteAddress=" + querylog.CRLF +
		"Program=" + querylog.UNKNOWN_APPLICATION + querylog.CRLF +
		"History=" + querylog.CRLF +
		"EOE"
}