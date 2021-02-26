package log

import (
	"fmt"
	"strings"

	slog "github.com/souliot/siot-log"
)

var DefaultBeeLogger = newDefaultBeeLogger()

type BeeLogger struct {
	log    *slog.BeeLogger
	fields Fields
}

func NewBeeLogger(beelog *slog.BeeLogger, fields Fields) Logger {
	return &BeeLogger{
		log:    beelog,
		fields: fields,
	}
}

func newDefaultBeeLogger() Logger {
	logger := slog.NewLogger()
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)
	logger.EnableFullFilePath(true)
	logger.SetLevel(slog.LevelInfo)

	return NewBeeLogger(logger, nil)
}

func (l *BeeLogger) Trace(f interface{}, v ...interface{}) {
	l.log.Trace(formatLog(f, v...))
}

func (l *BeeLogger) Debug(f interface{}, v ...interface{}) {
	l.log.Debug(formatLog(f, v...))
}

func (l *BeeLogger) Info(f interface{}, v ...interface{}) {
	l.log.Info(formatLog(f, v...))
}

func (l *BeeLogger) Warn(f interface{}, v ...interface{}) {
	l.log.Warn(formatLog(f, v...))
}

func (l *BeeLogger) Error(f interface{}, v ...interface{}) {
	l.log.Error(formatLog(f, v...))
}

func (l *BeeLogger) Fatal(f interface{}, v ...interface{}) {
	l.log.Critical(formatLog(f, v...))
}

func (l *BeeLogger) Panic(f interface{}, v ...interface{}) {
	l.log.Emergency(formatLog(f, v...))
}

func (l *BeeLogger) WithPrefix(prefix string) Logger {
	l.log.SetPrefix(prefix)
	return NewBeeLogger(l.log, l.GetFields())
}

func (l *BeeLogger) WithFields(fields Fields) Logger {
	return NewBeeLogger(l.log, l.GetFields().WithFields(fields))
}

func (l *BeeLogger) GetFields() Fields {
	return l.fields
}

func (l *BeeLogger) GetPrefix() string {
	return l.log.GetPrefix()
}

func (l *BeeLogger) SetLevel(lv int) {
	l.log.SetLevel(lv)
}

func (l *BeeLogger) GetLevel() int {
	return l.log.GetLevel()
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if !strings.Contains(msg, "%") {
			// do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
