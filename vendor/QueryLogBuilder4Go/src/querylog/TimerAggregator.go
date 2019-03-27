// Copyright 2018 The Go Authors. All rights reserved.
// Author: jeremyyin (Dongxun Yin)
package querylog

import (
	"github.com/juju/errors"
	"QueryLogBuilder4Go/src/common"
	"bytes"
	"strconv"
)

const (
	// "-open-" is added to timer names if they are serialized while still running
	OPEN_MARKER string = "-open-"
)
type TimerAggregator struct {
	// TODO: Add logger
	name 			string
	currentCount 	int64
	lastTime 		int64
	totalCount 		int64
	totalElapsed 	int64
	isFinished 		bool
}

func NewTimerAggregator(name string) (*TimerAggregator, error) {
	if len(name) == 0 {
		return nil, errors.New("name can not be empty")
	}

	agg := TimerAggregator{}

	agg.name = name
	agg.currentCount = 0
	agg.lastTime = 0
	agg.totalCount = 0
	agg.totalElapsed = 0
	agg.isFinished = false

	return &agg, nil
}

func (agg *TimerAggregator) GetName() string {
	return agg.name
}

// start another concurrent event
// pass NOW as param since we do not want to call timer.Now() too often
func (agg *TimerAggregator) Start(now int64) {
	if agg.currentCount > 0 {
		agg.totalElapsed += agg.currentCount * (now - agg.lastTime)
	}

	agg.lastTime = now
	agg.totalCount++
	agg.currentCount++
}

func (agg *TimerAggregator) End(now int64) {
	if agg.currentCount < 1 {
		return
	}

	agg.totalElapsed += agg.currentCount * (now - agg.lastTime)
	agg.lastTime = now
	agg.currentCount--
}

func (agg *TimerAggregator) AddTime(elapseTime int64) {
	agg.AddTimeWithSample(elapseTime, 1)
}

func (agg *TimerAggregator) AddTimeWithSample(elapseTime int64, numSample int64) {
	agg.totalCount += numSample
	agg.totalElapsed += elapseTime
}

func (agg *TimerAggregator) Finish(timeSource querylog.TimeSource) {
	agg.isFinished = true

	if agg.currentCount == 0 {
		return
	}

	now := timeSource.CurrentTimeMillis()

	agg.totalElapsed += agg.currentCount * (now - agg.lastTime)
	agg.lastTime = now
	agg.currentCount = 0
}

func (agg *TimerAggregator) ToStringWithTimeSource(timeSource querylog.TimeSource) (string, error) {
	if agg.currentCount == 0 {
		return agg.ToString()
	}

	now := timeSource.CurrentTimeMillis()
	elapsed := agg.totalElapsed + agg.currentCount * (now - agg.lastTime)

	var builder bytes.Buffer
	// WriteString will return errTooLarge error if the buffer is too large
	// But we will assume this will never happen in query log builder
	// So just ignore it
	builder.WriteString(agg.name)
	builder.WriteString(OPEN_MARKER)

	builder.WriteString(strconv.FormatInt(agg.currentCount, 10))
	builder.WriteByte(':')
	builder.WriteString(strconv.FormatInt(elapsed, 10))
	builder.WriteByte('/')
	builder.WriteString(strconv.FormatInt(agg.totalCount, 10))

	return builder.String(), nil
}

func (agg *TimerAggregator) ToString() (string, error) {
	if agg.currentCount > 0 {
		return "", errors.New("cannot call ToString() with open timers")
	}

	var builder bytes.Buffer

	builder.WriteString(agg.name)
	builder.WriteByte(':')
	builder.WriteString(strconv.FormatInt(agg.totalElapsed, 10))
	builder.WriteByte('/')
	builder.WriteString(strconv.FormatInt(agg.totalCount, 10))

	return builder.String(), nil
}

func (agg *TimerAggregator) GetTotalCount() int64 {
	return agg.totalCount
}

func (agg *TimerAggregator) GetTotalElapsed() int64 {
	return agg.totalElapsed
}