// Package logger
//
// @author: xwc1125
package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Logger interface {
	Name() string
	New(module string, ctx ...interface{}) Logger

	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})

	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

var (
	root        Logger
	logAdapters = make([]Logger, 0)
	lock        sync.RWMutex
)

// RegisterLog 需要先注册logger
func RegisterLog(logger Logger) Logger {
	if logger == nil {
		return logger
	}
	lock.Lock()
	defer lock.Unlock()
	logAdapters = append(logAdapters, logger)
	root = logger
	return logger
}

// SwitchLog 切换日志kit
func SwitchLog(logKitName string) {
	lock.RLock()
	defer lock.RUnlock()
	for _, logger := range logAdapters {
		if strings.EqualFold(logger.Name(), logKitName) {
			root = logger
		}
	}
}

// New 方法用到root，因此此方法不能够应用到init()中，除非先执行了RegisterLog
func New(module string, ctx ...interface{}) Logger {
	return root.New(module, ctx...)
}

func Trace(msg string, ctx ...interface{}) {
	if root == nil {
		return
	}
	root.Trace(msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	if root == nil {
		return
	}
	root.Debug(msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	if root == nil {
		return
	}
	root.Info(msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	if root == nil {
		return
	}
	root.Warn(msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	if root == nil {
		fmt.Println(msg, ctx)
		return
	}
	root.Error(msg, ctx...)
}

func Crit(msg string, ctx ...interface{}) {
	if root == nil {
		fmt.Println(msg, ctx)
		os.Exit(1)
		return
	}
	root.Crit(msg, ctx...)
}

func Printf(format string, v ...interface{}) {
	if root == nil {
		fmt.Printf(format, v...)
		return
	}
	root.Printf(format, v...)
}

func Print(v ...interface{}) {
	if root == nil {
		fmt.Print(v...)
		return
	}
	root.Print(v...)
}

func Println(v ...interface{}) {
	if root == nil {
		fmt.Println(v...)
		return
	}
	root.Println(v...)
}

func Fatal(v ...interface{}) {
	if root == nil {
		fmt.Print(v...)
		os.Exit(1)
		return
	}
	root.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	if root == nil {
		fmt.Printf(format, v...)
		os.Exit(1)
		return
	}
	root.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	if root == nil {
		fmt.Println(v...)
		os.Exit(1)
		return
	}
	root.Fatalln(v...)
}

func Panic(v ...interface{}) {
	if root == nil {
		fmt.Print(v...)
		os.Exit(1)
		return
	}
	root.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	if root == nil {
		fmt.Printf(format, v...)
		os.Exit(1)
		return
	}
	root.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	if root == nil {
		fmt.Println(v...)
		os.Exit(1)
		return
	}
	root.Panicln(v)
}
