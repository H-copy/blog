package logger

import (
	"fmt"
	"testing"
)

func TestLevelString(t *testing.T) {

	levelList := []Level{
		LevelDebug,
		LevelInfo,
		LevelWarn,
		LevelError,
		LevelFatal,
		LevelPanic,
	}

	for _, level := range levelList {

		txt := level.String()
		fmt.Printf("code: %d, txt: %s\n", level, txt)

		if txt == "" {
			t.Fatal("level txt is undefined: ", level)
		}

	}

}
