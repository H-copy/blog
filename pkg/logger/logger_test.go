package logger

import (
	"fmt"
	"testing"
	"os"
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


func TestWidth(t *testing.T){

	file, err := os.OpenFile("./test.log", os.O_RDWR, 0666 )

	if err != nil{
		fmt.Println("err: ", err)
		return
	}

	defer file.Close()
	
	
	l := NewLogger(file, "log_", 0)

	if nl := l.WithLevel(LevelInfo); nl.level != LevelInfo {
		t.Fatal("withLevel fail", nl.level, LevelInfo)
	}

	f := Fields{"zero": 0, "one":1}
	
	if nl := l.WithFields(f); nl.fields["zero"] != f["zero"]{
		t.Fatal("withFields fail", nl.fields["zero"], f["zero"])
	}
	
}