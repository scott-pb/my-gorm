package log

import (
	"io"
	"log"
	"os"
	"sync"
)

type LevelFlag uint

const (
	DebugLevel LevelFlag = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DisableLevel
)

// 颜色设置
var debugLog = log.New(os.Stdout, "\033[1;37;40m[mygorm-debug]", log.Ldate|log.Lmicroseconds|log.Lshortfile)
var infoLog = log.New(os.Stdout, "\033[1;32;40m[mygorm-info]", log.Ldate|log.Lmicroseconds|log.Lshortfile)
var warnLog = log.New(os.Stdout, "\033[1;33;40m[mygorm-warn]", log.Ldate|log.Lmicroseconds|log.Lshortfile)
var errorLog = log.New(os.Stdout, "\033[1;31;40m[mygorm-error]", log.Ldate|log.Lmicroseconds|log.Lshortfile)

var loggers = []*log.Logger{debugLog, infoLog, warnLog, errorLog}
var mu sync.Mutex

// 日志方法
var (
	Debug  = debugLog.Println
	DebugF = debugLog.Printf

	Info  = infoLog.Println
	InfoF = infoLog.Printf

	Warn  = warnLog.Println
	WarnF = warnLog.Printf

	Error  = errorLog.Println
	ErrorF = errorLog.Printf
)

// SetLogLevel 设置日志级别
func SetLogLevel(level LevelFlag) {
	//加锁
	mu.Lock()
	//释放锁
	defer mu.Unlock()

	//先把所有的日志级别允许输出
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	//设置的level 大于err 说明err以下的都不输出
	if level > ErrorLevel {
		errorLog.SetOutput(io.Discard)
	}

	//只有 level 大于 warn才会输出，只有err 级别才会输出
	if level > WarnLevel {
		warnLog.SetOutput(io.Discard)
	}

	if level > InfoLevel {
		infoLog.SetOutput(io.Discard)
	}

	if level > DebugLevel {
		debugLog.SetOutput(io.Discard)
	}

}
