package querylog

import (
	"sync"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)

// An implementation of event data formatter
// We do not need any pre defined values in logger
type EventDataLogFormatter struct {
	// Whether the logger's out is to a terminal
	isTerminal bool

	sync.Once
}

func (f *EventDataLogFormatter) init(entry *logrus.Entry) {
	if entry.Logger != nil {
		f.isTerminal = checkIfTerminal(entry.Logger.Out)
	}
}

// Format renders a single log entry
func (f *EventDataLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.Do(func() { f.init(entry) })

	f.appendValue(b, entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *EventDataLogFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}


