package logger

import (

	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)


type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var (
	LogLevelTxt = map[Level]string{
		LevelDebug: "debug",
		LevelInfo:  "info",
		LevelWarn:  "warn",
		LevelError: "error",
		LevelFatal: "fatal",
		LevelPanic: "panic",
	}
)

func (l Level) String() string {
	if txt, ok := LogLevelTxt[l]; ok {
		return txt
	} else {
		return ""
	}
}


type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
	Level     Level
}



func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) Clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(level Level) *Logger {
	nl := l.Clone()
	nl.level = level
	return nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	nl := l.Clone()

	if nl.fields == nil {
		nl.fields = make(Fields)
	}

	for k, v := range f {
		nl.fields[k] = v
	}

	return nl
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := l.Clone()

	nl.ctx = ctx

	return nl
}

func (l *Logger) WithCaller(skip int) *Logger {
	nl := l.Clone()
	pc, file, line, ok := runtime.Caller(skip)

	if ok {
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return nl
}

func (l *Logger) WithCallersFrames() *Logger {

	maxCallerDepth := 25
	minCallerDepth := 1

	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))

		if !more {
			break
		}
	}

	nl := l.Clone()
	nl.callers = callers
	return nl

}




func (l *Logger) JSONForamt(message string) map[string]interface{} {

	data := make(Fields, len(l.fields))

	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.fields) > 0 {

		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}

	}

	return data
}

func (l *Logger) Output(message string) {

	body, _ := json.Marshal(l.JSONForamt(message))

	content := string(body)

	switch l.level {
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	default:
		l.newLogger.Println(content)
	}

}

func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf( format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}
