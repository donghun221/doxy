package querylog

import "sync"

type ThreadSafeEventDataImpl struct {
	delegate 	*EventDataImpl
	lock 		*sync.Mutex
}

func NewThreadSafeEventDataImpl(delegate *EventDataImpl) *ThreadSafeEventDataImpl {
	event := ThreadSafeEventDataImpl{}
	event.delegate = delegate
	event.lock = &sync.Mutex{}
	return &event
}

func (event *ThreadSafeEventDataImpl) GetApplicationName() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetApplicationName()
}

func (event *ThreadSafeEventDataImpl) GetHostName() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetHostName()
}

func (event *ThreadSafeEventDataImpl) GetOperation() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetOperation()
}

func (event *ThreadSafeEventDataImpl) SetOperation(operation string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetOperation(operation)
}

func (event *ThreadSafeEventDataImpl) GetStatus() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetStatus()
}

func (event *ThreadSafeEventDataImpl) SetStatus(status string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetStatus(status)
}

func (event *ThreadSafeEventDataImpl) GetQueryLogStatus() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetQueryLogStatus()
}

func (event *ThreadSafeEventDataImpl) SetQueryLogStatus(status string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetQueryLogStatus(status)
}

func (event *ThreadSafeEventDataImpl) Reset() {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.Reset()
}

func (event *ThreadSafeEventDataImpl) GetStartTime() int64 {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetStartTime()
}

func (event *ThreadSafeEventDataImpl) SetStartTime(startTime int64) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetStartTime(startTime)
}

func (event *ThreadSafeEventDataImpl) GetEndTime() int64 {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetEndTime()
}

func (event *ThreadSafeEventDataImpl) SetEndTime(endTime int64) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetEndTime(endTime)
}

func (event *ThreadSafeEventDataImpl) StartTimer(timerName string) error {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.StartTimer(timerName)
}

func (event *ThreadSafeEventDataImpl) EndTimer(name string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.EndTimer(name)
}

func (event *ThreadSafeEventDataImpl) UpdateTimer(name string, elapsed int64) error {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.UpdateTimer(name, elapsed)
}

func (event *ThreadSafeEventDataImpl) UpdateTimerWithSample(name string, elapsed, sample int64) error {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.UpdateTimerWithSample(name, elapsed, sample)
}

func (event *ThreadSafeEventDataImpl) GetTimeElapsed(timerName string) int64 {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetTimeElapsed(timerName)
}

func (event *ThreadSafeEventDataImpl) GetRemoteAddr() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetRemoteAddr()
}

func (event *ThreadSafeEventDataImpl) SetRemoteAddr(addr string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetRemoteAddr(addr)
}

func (event *ThreadSafeEventDataImpl) GetCounter(name string) int64 {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetCounter(name)
}

func (event *ThreadSafeEventDataImpl) SetCounter(name string, value int64) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.SetCounter(name, value)
}

func (event *ThreadSafeEventDataImpl) InCCounter(name string, value int64) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.InCCounter(name, value)
}

func (event *ThreadSafeEventDataImpl) AddNameValuePair(name, value string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.AddNameValuePair(name, value)
}

func (event *ThreadSafeEventDataImpl) AppendNameValuePair(name, value string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.AppendNameValuePair(name, value)
}

func (event *ThreadSafeEventDataImpl) GetValue(name string) string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.GetValue(name)
}

func (event *ThreadSafeEventDataImpl) FinishCurrentEvent(name string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.FinishCurrentEvent(name)
}

func (event *ThreadSafeEventDataImpl) RecordHistoryEvent(name string) {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.RecordHistoryEvent(name)
}

func (event *ThreadSafeEventDataImpl) RecordProfiledData() {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.RecordProfiledData()
}

func (event *ThreadSafeEventDataImpl) ToJsonFormat() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.ToJsonFormat()
}

func (event *ThreadSafeEventDataImpl) ToEventDataFormat() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.ToEventDataFormat()
}

func (event *ThreadSafeEventDataImpl) recordProfiledData() {
	event.lock.Lock()
	defer event.lock.Unlock()

	event.delegate.RecordProfiledData()
}

func (event *ThreadSafeEventDataImpl) toString() string {
	event.lock.Lock()
	defer event.lock.Unlock()

	return event.delegate.toString()
}

func (event *ThreadSafeEventDataImpl) getCount(counterName string) int64 {
	event.lock.Lock()
	event.lock.Unlock()

	return event.delegate.getCount(counterName)
}