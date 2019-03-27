package tst

import (
	"testing"
	"QueryLogBuilder4Go/src/querylog"
	"github.com/stretchr/testify/assert"
	"QueryLogBuilder4Go/tst/base"
)

func TestTimerAggregator_WithEmptyName(t *testing.T) {
	agg, err := querylog.NewTimerAggregator("")

	assert.NotNil(t, err)
	assert.Nil(t, agg)
}

func TestTimerAggregator_WithSingleEvent(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	agg.End(2 * tst.TEN)

	assert.Equal(t, tst.ONE, agg.GetTotalCount())
	assert.Equal(t,tst.TEN, agg.GetTotalElapsed())
}

func TestTimerAggregator_WithMultipleEvent(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	agg.Start(tst.TEN + tst.FIVE)
	// elapse by two(17 - 15)
	agg.End(tst.TEN + tst.SEVEN)
	// elapse by ten(20 - 10)
	agg.End(tst.TEN + tst.TEN)

	// It is obvious to see there were two event happened
	assert.Equal(t, tst.TWO, agg.GetTotalCount())

	// The time aggregator was designed to track each start-end event by stack like mechanism
	assert.Equal(t,tst.TEN + tst.TWO, agg.GetTotalElapsed())
}

func TestTimerAggregator_AddTimer_WithoutEnding(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.AddTime(tst.TEN)
	agg.AddTime(tst.TWO)

	// It is obvious to see there were two event happened
	assert.Equal(t, tst.TWO, agg.GetTotalCount())

	// The time aggregator was designed to track each start-end event by stack like mechanism
	assert.Equal(t,tst.TEN + tst.TWO, agg.GetTotalElapsed())
}

func TestTimerAggregator_AddTimer_WithEnding(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	agg.Start(tst.TEN + tst.FIVE)
	// elapse by ten
	agg.AddTime(tst.TEN)
	// elapse by two
	agg.AddTime(tst.TWO)
	// elapse by two(17 - 15)
	agg.End(tst.TEN + tst.SEVEN)
	// elapse by 10(20 - 10)
	agg.End(tst.TEN + tst.TEN)

	// It is obvious to see there were four time event happened
	assert.Equal(t, tst.FOUR, agg.GetTotalCount())

	// The time aggregator was designed to track each start-end event by stack like mechanism
	assert.Equal(t,tst.TEN + tst.TEN + tst.FOUR, agg.GetTotalElapsed())
}

func TestTimerAggregator_WithTimerLeftOpen_One(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	agg.Start(tst.TEN + tst.FIVE)

	timeSource := tst.TestTimeSource{20}

	res, err := agg.ToStringWithTimeSource(&timeSource)
	assert.Nil(t, err)
	assert.Equal(t, tst.MOCK_EVENT_NAME + "-open-2:15/2", res)
}

func TestTimerAggregator_WithTimerLeftOpen_Two(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	agg.Start(tst.TEN + tst.TWO)
	agg.End(tst.TEN + tst.THREE)
	agg.Start(tst.TEN + tst.FIVE)

	timeSource := tst.TestTimeSource{20}

	res, err := agg.ToStringWithTimeSource(&timeSource)
	assert.Nil(t, err)
	assert.Equal(t, tst.MOCK_EVENT_NAME + "-open-2:16/3", res)
}

func TestTimerAggregator_Incremental(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.AddTimeWithSample(1444, 10)

	res, err := agg.ToString()
	assert.Nil(t, err)
	assert.Equal(t, tst.MOCK_EVENT_NAME + ":1444/10", res)
}

func TestTimerAggregator_ToStringWithOpenTimer(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.Start(tst.TEN)
	_, err := agg.ToString()
	assert.NotNil(t, err)
}

func TestTimerAggregator_ToString_HappyCase(t *testing.T) {
	agg := tst.GetTimerAggregator(t)

	agg.AddTime(tst.TEN)
	agg.AddTime(tst.TWO)

	res, err := agg.ToString()
	assert.Nil(t, err)
	assert.Equal(t, tst.MOCK_EVENT_NAME + ":12/2", res)
}