package log

import (
	"fmt"
	"strings"
)

// Logger interface used as base logger throughout the library.
type Logger interface {
	Trace(f interface{}, v ...interface{})

	Debug(f interface{}, v ...interface{})

	Info(f interface{}, v ...interface{})

	Warn(f interface{}, v ...interface{})

	Error(f interface{}, v ...interface{})

	Fatal(f interface{}, v ...interface{})

	Panic(f interface{}, v ...interface{})

	WithPrefix(prefix string) Logger
	GetPrefix() string

	WithFields(fields Fields) Logger
	GetFields() Fields

	SetLevel(int)
	GetLevel() int
}

type Loggable interface {
	Log() Logger
}

type Fields map[string]interface{}

func (fields Fields) String() string {
	str := make([]string, 0)

	for k, v := range fields {
		str = append(str, fmt.Sprintf("%s=%+v", k, v))
	}

	return strings.Join(str, " ")
}

func AddFieldsFrom(logger Logger, values ...interface{}) Logger {
	for _, value := range values {
		switch v := value.(type) {
		case Logger:
			logger = logger.WithFields(v.GetFields())
		case Loggable:
			logger = logger.WithFields(v.Log().GetFields())
		case interface{ Fields() Fields }:
			logger = logger.WithFields(v.Fields())
		}
	}
	return logger
}

func (fields Fields) WithFields(newFields Fields) Fields {
	allFields := make(Fields)

	for k, v := range fields {
		allFields[k] = v
	}

	for k, v := range newFields {
		allFields[k] = v
	}

	return allFields
}
