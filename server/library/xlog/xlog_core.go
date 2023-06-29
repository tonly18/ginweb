package xlog

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"sync"
	"time"
)

// Log Level
const (
	LogDebug = iota
	LogInfo
	LogWarn
	LogError
	LogPanic
	LogFatal
)

// Log Level String
var levels = []string{
	"[DEBUG]",
	"[INFO]",
	"[WARN]",
	"[ERROR]",
	"[PANIC]",
	"[FATAL]",
}

type Logger struct {
	file   *os.File
	buffer bytes.Buffer
	writer *Writer
	mu     sync.Mutex

	calldDepth int
}

func NewLogger() *Logger {
	writer := &Logger{
		writer:     NewWriter(bufio.NewWriter(os.Stderr)),
		mu:         sync.Mutex{},
		calldDepth: 2,
	}

	runtime.SetFinalizer(writer, CleanLogger)

	return writer
}

func CleanLogger(logger *Logger) {
	if logger.file != nil {
		logger.file.Close()
	}
}

//SetLogFile 设置日志文件
func (logger *Logger) SetLogFile(file *os.File) {
	if logger.file != nil {
		logger.file.Close()
	}
	logger.file = file
	logger.writer = NewWriter(bufio.NewWriter(file))
}

func (logger *Logger) Log(level int, data []byte) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	now := time.Now()
	var file string
	var line int
	_, file, line, ok := runtime.Caller(logger.calldDepth)
	if !ok {
		file = "unknown-file"
		line = 0
	}

	logger.buffer.Reset()
	logger.formatHeader(now, file, line, level)
	logger.buffer.Write(data)
	logger.writer.Write(logger.buffer.Bytes())
	logger.writer.Write([]byte("\n"))
}

func (logger *Logger) formatHeader(t time.Time, file string, line int, level int) {
	buf := &logger.buffer

	//日期时间
	year, month, day := t.Date()
	itoa(buf, year, 4)
	buf.WriteByte('/') // "2019/"
	itoa(buf, int(month), 2)
	buf.WriteByte('/') // "2019/04/"
	itoa(buf, day, 2)
	buf.WriteByte(' ') // "2019/04/11 "

	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	buf.WriteByte(':') // "11:"
	itoa(buf, min, 2)
	buf.WriteByte(':') // "11:15:"
	itoa(buf, sec, 2)  // "11:15:33"
	// Microsecond flag is set
	buf.WriteByte(' ')

	//日志类型
	buf.Write([]byte(levels[level]))

	//文件名
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			// Get the file name after the last '/' character, e.g. "zinx.go" from "/home/go/src/zinx.go"
			short = file[i+1:]
			break
		}
	}
	file = short

	buf.Write([]byte(file))
	buf.WriteByte(':')
	itoa(buf, line, -1) // line number
	buf.Write([]byte(" "))
}

// 将一个整形转换成一个固定长度的字符串,字符串宽度应该是大于0的要确保buffer是有容量空间的
func itoa(buf *bytes.Buffer, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		buf.WriteByte('0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}

	// avoID slicing b to avoID an allocation.
	for bp < len(b) {
		buf.WriteByte(b[bp])
		bp++
	}
}
