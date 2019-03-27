package tst

import (
	"testing"
	"QueryLogBuilder4Go/tst/base"
	"github.com/stretchr/testify/assert"
	"QueryLogBuilder4Go/src/common"
)

func TestNoopEventData_GetApplicationName(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetApplicationName())
}

func TestNoopEventData_GetHostName(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetHostName())
}

func TestNoopEventData_GetOperation(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetOperation())
}

func TestNoopEventData_SetOperation(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetOperation(querylog.EMPTY_STRING)
}

func TestNoopEventData_GetStatus(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetStatus())
}

func TestNoopEventData_SetStatus(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetStatus(querylog.EMPTY_STRING)
}

func TestNoopEventData_GetQueryLogStatus(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetQueryLogStatus())
}

func TestNoopEventData_SetQueryLogStatus(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetQueryLogStatus(querylog.EMPTY_STRING)
}

func TestNoopEventData_Reset(t *testing.T) {
	event := tst.GetNoopEventData()
	event.Reset()
}

func TestNoopEventData_GetStartTime(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Zero(t, event.GetStartTime())
}

func TestNoopEventData_SetStartTime(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetStartTime(0)
}

func TestNoopEventData_GetEndTime(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Zero(t, event.GetEndTime())
}

func TestNoopEventData_SetEndTime(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetEndTime(0)
}

func TestNoopEventData_StartTimer(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Nil(t, event.StartTimer(querylog.EMPTY_STRING))
}

func TestNoopEventData_EndTimer(t *testing.T) {
	event := tst.GetNoopEventData()
	event.EndTimer(querylog.EMPTY_STRING)
}

func TestNoopEventData_UpdateTimer(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Nil(t, event.UpdateTimer(querylog.EMPTY_STRING, tst.ZERO))
}

func TestNoopEventData_UpdateTimerWithSample(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Nil(t, event.UpdateTimerWithSample(querylog.EMPTY_STRING, tst.ZERO, tst.ZERO))
}

func TestNoopEventData_GetTimeElapsed(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Zero(t, event.GetTimeElapsed(querylog.EMPTY_STRING))
}

func TestNoopEventData_GetRemoteAddr(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetRemoteAddr())
}

func TestNoopEventData_SetRemoteAddr(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetRemoteAddr(querylog.EMPTY_STRING)
}

func TestNoopEventData_GetCounter(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Zero(t, event.GetCounter(querylog.EMPTY_STRING))
}

func TestNoopEventData_SetCounter(t *testing.T) {
	event := tst.GetNoopEventData()
	event.SetCounter(querylog.EMPTY_STRING, tst.ZERO)
}

func TestNoopEventData_InCCounter(t *testing.T) {
	event := tst.GetNoopEventData()
	event.InCCounter(querylog.EMPTY_STRING, tst.ZERO)
}

func TestNoopEventData_AddNameValuePair(t *testing.T) {
	event := tst.GetNoopEventData()
	event.AddNameValuePair(querylog.EMPTY_STRING, querylog.EMPTY_STRING)
}

func TestNoopEventData_AppendNameValuePair(t *testing.T) {
	event := tst.GetNoopEventData()
	event.AppendNameValuePair(querylog.EMPTY_STRING, querylog.EMPTY_STRING)
}

func TestNoopEventData_GetValue(t *testing.T) {
	event := tst.GetNoopEventData()
	assert.Empty(t, event.GetValue(querylog.EMPTY_STRING))
}

func TestNoopEventData_FinishCurrentEvent(t *testing.T) {
	event := tst.GetNoopEventData()
	event.FinishCurrentEvent(querylog.EMPTY_STRING)
}

func TestNoopEventData_RecordHistoryEvent(t *testing.T) {
	event := tst.GetNoopEventData()
	event.RecordHistoryEvent(querylog.EMPTY_STRING)
}

func TestNoopEventData_RecordProfiledData(t *testing.T) {
	event := tst.GetNoopEventData()
	event.RecordProfiledData()
}

func TestNoopEventData_ToJsonFormat(t *testing.T) {
	event := tst.GetNoopEventData()
	event.ToJsonFormat()
}

func TestNoopEventData_ToEventDataFormat(t *testing.T) {
	event := tst.GetNoopEventData()
	event.ToEventDataFormat()
}