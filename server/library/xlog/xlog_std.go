package xlog

import (
	"bytes"
	"fmt"
	"os"
)

var stdXLog *Logger = NewLogger()

func SetLogFile(file *os.File) {
	stdXLog.SetLogFile(file)
}

func Debug(data ...any) {
	args := bytes.Buffer{}
	for _, v := range data {
		args.Write([]byte(fmt.Sprintf(`%v`, v)))
	}
	stdXLog.Log(LogDebug, args.Bytes())
}

func Debugf(format string, data ...any) {
	args := fmt.Sprintf(format, data...)
	stdXLog.Log(LogDebug, []byte(args))
}

func Info(data ...any) {
	args := bytes.Buffer{}
	for _, v := range data {
		args.Write([]byte(fmt.Sprintf(`%v`, v)))
	}
	stdXLog.Log(LogInfo, args.Bytes())
}

func Infof(format string, data ...any) {
	args := fmt.Sprintf(format, data...)
	stdXLog.Log(LogInfo, []byte(args))
}

func Error(data ...any) {
	args := bytes.Buffer{}
	for _, v := range data {
		args.Write([]byte(fmt.Sprintf(`%v`, v)))
	}
	stdXLog.Log(LogError, args.Bytes())
}

func Errorf(format string, data ...any) {
	args := fmt.Sprintf(format, data...)
	stdXLog.Log(LogError, []byte(args))
}
