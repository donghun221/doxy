// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

import "QueryLogBuilder4Go/src/common"

type NoopEventData struct {}

func (NoopEventData) GetApplicationName() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) GetHostName() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) GetOperation() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) SetOperation(string) {
	// No op
}

func (NoopEventData) GetStatus() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) SetStatus(string) {
	// No op
}

func (NoopEventData) GetQueryLogStatus() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) SetQueryLogStatus(string) {
	// No op
}

func (NoopEventData) Reset() {
	// No op
}

func (NoopEventData) GetStartTime() int64 {
	return 0
}

func (NoopEventData) SetStartTime(int64) {
	// No op
}

func (NoopEventData) GetEndTime() int64 {
	return 0
}

func (NoopEventData) SetEndTime(endTime int64) {
	// No op
}

func (NoopEventData) StartTimer(string) error {
	return nil
}

func (NoopEventData) EndTimer(string) {
	// No op
}

func (NoopEventData) UpdateTimer(string, int64) error {
	return nil
}

func (NoopEventData) UpdateTimerWithSample(string, int64, int64) error {
	return nil
}

func (NoopEventData) GetTimeElapsed(string) int64 {
	return 0
}

func (NoopEventData) GetRemoteAddr() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) SetRemoteAddr(string) {
	// No op
}

func (NoopEventData) GetCounter(string) int64 {
	return 0
}

func (NoopEventData) SetCounter(string, int64) {
	// No op
}

func (NoopEventData) InCCounter(string, int64) {
	// No op
}

func (NoopEventData) AddNameValuePair(name, value string) {
	// No op
}

func (NoopEventData) AppendNameValuePair(name, value string) {
	// No op
}

func (NoopEventData) GetValue(name string) string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) FinishCurrentEvent(name string) {
	// No op
}

func (NoopEventData) RecordHistoryEvent(name string) {
	// No op
}

func (NoopEventData) RecordProfiledData() {
	// No op
}

func (NoopEventData) ToJsonFormat() string {
	return querylog.EMPTY_STRING
}

func (NoopEventData) ToEventDataFormat() string {
	return querylog.EMPTY_STRING
}


