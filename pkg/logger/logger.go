package logger

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
