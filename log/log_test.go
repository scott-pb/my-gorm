package log

import (
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	//设置级别为err
	SetLogLevel(ErrorLevel)
	Error("error TestSetLogLevel") //会输出
	Info("info TestSetLogLevel")   //不会输出

}

func TestSetLogLevel2(t *testing.T) {
	//设置级别为info
	SetLogLevel(InfoLevel)
	Error("error TestSetLogLevel2") //会输出
	Info("info TestSetLogLevel")    //会输出

}
