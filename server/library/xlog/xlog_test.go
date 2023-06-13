package xlog

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestXlog(t *testing.T) {
	file, _ := os.OpenFile("./xlog.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	//logger := NewLogger(file)
	//logger.CloseLogFile(file)
	//err := logger.SetLogFile("./log/xlog.log")
	//fmt.Println("err::::::", err)

	SetLogFile(file)

	for i := 0; i < 5; i++ {
		//logger.Log(LogDebug, []byte(fmt.Sprintf(`abc:%v`, i)))
		//go Debug("abc-123-def: ", i)
		Debug("abc-123-def: ", i)
		//time.Sleep(1 * time.Second)
	}

	time.Sleep(12 * time.Second)
	fmt.Println("main")
}
