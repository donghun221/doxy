package querylog

import (
	"QueryLogBuilder4Go/src/common"
	"github.com/sirupsen/logrus"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
)

const (
	UNKNOWN_APPLICATION string = "** Unkown Application **"
	UNKNOWN_HOST_NAME string = "** Unknown Host Name **"
)

type EventDataFactory struct {
	timeSource 			querylog.TimeSource
	applicationName 	string
	hostName 			string
	queryLog 			*logrus.Logger
	listeners			[]QueryLogEntryListener
	defaultNameValues	map[string]string
	lock				*sync.Mutex
}

func NewEventDataFactoryDefault() *EventDataFactory {
	return NewEventDataFactory(nil, querylog.EMPTY_STRING, querylog.EMPTY_STRING, nil, nil)
}

func NewEventDataFactory(
	timeSource querylog.TimeSource,
	applicationName,
	hostName string,
	queryLog *logrus.Logger,
	loggerConf *lumberjack.Logger) *EventDataFactory {
	factory := EventDataFactory{}

	if timeSource == nil {
		timeSource = &querylog.RealTimeSource{}
	}

	if len(applicationName) == 0 {
		applicationName = UNKNOWN_APPLICATION
	}

	if len(hostName) == 0 {
		hostName = ObtainHostName()
	}

	if queryLog == nil {
		queryLog = factory.initEventDataLogger(loggerConf)
	}

	factory.timeSource = timeSource
	factory.applicationName = applicationName
	factory.hostName = hostName
	factory.queryLog = queryLog

	factory.listeners = make([]QueryLogEntryListener, 0, 0)
	factory.defaultNameValues = make(map[string]string)

	return &factory
}

func (factory *EventDataFactory) CreateEventData() EventData {
	event := NewEventDataImpl(
		factory.GetTimeSource(),
		factory.GetApplicationName(),
		factory.GetHostName(),
		factory.defaultNameValues,
		factory.queryLog,
		factory.listeners,
		false)

	return event
}

func (factory *EventDataFactory) CreateThreadSafeEventData() EventData {
	event := NewEventDataImpl(
		factory.GetTimeSource(),
		factory.GetApplicationName(),
		factory.GetHostName(),
		factory.defaultNameValues,
		factory.queryLog,
		factory.listeners,
		false)

	return NewThreadSafeEventDataImpl(event)
}

func (factory *EventDataFactory) CreateNoopEventData() EventData {
	event := NoopEventData{}
	return &event
}

func (factory *EventDataFactory) initEventDataLogger(loggerConf *lumberjack.Logger) *logrus.Logger {
	logger := logrus.New()

	// Add formatter
	formatter := &EventDataLogFormatter{}
	logger.Formatter = formatter

	if loggerConf == nil {
		loggerConf = &lumberjack.Logger{
			Filename:   querylog.EVENT_DATA_LOG_PATH,
			MaxSize:    querylog.EVENT_DATA_MAX_LOG_SIZE, // 1GB by default
			MaxBackups: querylog.EVENT_DATA_MAX_BACKUPS,
			MaxAge:     querylog.EVENT_DATA_MAX_LOG_SIZE,
			LocalTime:  querylog.EVENT_DATA_LOCAL_TIME,
		}
	}

	logger.Out = loggerConf

	logger.Level = logrus.InfoLevel

	return logger
}

func (factory *EventDataFactory) GetApplicationName() string {
	return factory.applicationName
}

func (factory *EventDataFactory) GetListeners() []QueryLogEntryListener {
	return factory.listeners
}

func (factory *EventDataFactory) GetHostName() string {
	return factory.hostName
}

func (factory *EventDataFactory) GetTimeSource() querylog.TimeSource {
	return factory.timeSource
}

func (factory *EventDataFactory) AddDefaultNameValuePair(name, value string) {
	factory.lock.Lock()
	defer factory.lock.Unlock()

	factory.defaultNameValues[name] = value
}

func (factory *EventDataFactory) AddQueryLogEntryListener(listener QueryLogEntryListener) {
	factory.listeners = append(factory.listeners, listener)
}

func ObtainHostName() string {
	hostName, err := os.Hostname()

	// In this version, we will ignore errors returned by OS
	if err != nil {
		hostName = UNKNOWN_HOST_NAME
	}

	return hostName
}