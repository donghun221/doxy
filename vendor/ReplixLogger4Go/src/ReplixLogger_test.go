package replixlogger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"bytes"
)

type testLogSuite struct {
	buf *bytes.Buffer
}

var s testLogSuite = testLogSuite{&bytes.Buffer{}}

func TestStringToLogLevel(t *testing.T) {
	assert.Equal(t, log.FatalLevel, stringToLogLevel("fatal"))
	assert.Equal(t, log.ErrorLevel, stringToLogLevel("ERROR"))
	assert.Equal(t, log.WarnLevel, stringToLogLevel("warn"))
	assert.Equal(t, log.WarnLevel, stringToLogLevel("warning"))
	assert.Equal(t, log.DebugLevel, stringToLogLevel("debug"))
	assert.Equal(t, log.InfoLevel, stringToLogLevel("info"))
	assert.Equal(t, log.InfoLevel, stringToLogLevel("whatever"))
}

// TestLogging assure log format and log redirection works.
func TestLogging(t *testing.T) {
	conf := &LogConfig{Level: "warn", File: FileLogConfig{}}
	assert.Nil(t, InitLogger(conf))

	log.SetOutput(s.buf)

	log.Infof("[this message should not be sent to buf]")
	assert.Equal(t, 0, s.buf.Len())

	log.Warningf("[this message should be sent to buf]")
	entry, err := s.buf.ReadString('\n')
	assert.Nil(t, err)

	log.Warnf("this message comes from logrus")
	entry, err = s.buf.ReadString('\n')
	assert.Nil(t, err)
	assert.True(t, strings.Contains(entry, "log_test.go"))
}

func TestSlowQueryLogger(t *testing.T) {
	fileName := "slow_query"
	conf := &LogConfig{Level: "info", File: FileLogConfig{}, SlowQueryFile: fileName}
	err := InitLogger(conf)
	assert.Nil(t, err)

	defer os.Remove(fileName)

	SlowQueryLogger.Debug("debug message")
	SlowQueryLogger.Info("info message")
	SlowQueryLogger.Warn("warn message")
	SlowQueryLogger.Error("error message")
	assert.Equal(t, 0, s.buf.Len())

	f, err := os.Open(fileName)
	assert.Nil(t, err)

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		_, err = r.ReadString('\n')
		if err != nil {
			break
		}
	}

	assert.Equal(t, io.EOF, err)
}

func TestSlowQueryLoggerKeepOrder(t *testing.T) {
	fileName := "slow_query"
	conf := &LogConfig{Level: "warn", File: FileLogConfig{}, Format: "text", DisableTimestamp: true, SlowQueryFile: fileName}

	assert.Nil(t, InitLogger(conf))
	defer os.Remove(fileName)
	ft, ok := SlowQueryLogger.Formatter.(*textFormatter)

	assert.True(t, ok)

	assert.True(t, ft.EnableEntryOrder)

	SlowQueryLogger.Out = s.buf
	logEntry := log.NewEntry(SlowQueryLogger)
	logEntry.Data = log.Fields{
		"connectionId": 1,
		"costTime":     "1",
		"database":     "test",
		"sql":          "select 1",
		"txnStartTS":   1,
	}

	_, _, line, _ := runtime.Caller(0)
	logEntry.WithField("type", "slow-query").WithField("succ", true).Warnf("slow-query")
	expectMsg := fmt.Sprintf("log_test.go:%v: [warning] slow-query connectionId=1 costTime=1 database=test sql=select 1 succ=true txnStartTS=1 type=slow-query\n", line+1)

	assert.Equal(t, expectMsg, s.buf.String())

	s.buf.Reset()
	logEntry.Data = log.Fields{
		"a": "a",
		"d": "d",
		"e": "e",
		"b": "b",
		"f": "f",
		"c": "c",
	}

	_, _, line, _ = runtime.Caller(0)
	logEntry.Warnf("slow-query")
	expectMsg = fmt.Sprintf("log_test.go:%v: [warning] slow-query a=a b=b c=c d=d e=e f=f\n", line+1)

	assert.Equal(t, expectMsg, s.buf.String())
}
