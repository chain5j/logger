// Package test
//
// @author: xwc1125
package test

import (
	log "github.com/chain5j/log15"
	"github.com/chain5j/logger"
	"github.com/chain5j/logger/logrus"
	"github.com/chain5j/logger/zap"
	"testing"
)

func TestLog(t *testing.T) {
	zapLog()
	log.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>", "zapLog", "zapLog")
	//log15Log()
	log.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>", "log15Log", "log15Log")
	logrusLog()
	log.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>", "logrusLog", "logrusLog")
}

//func log15Log() {
//	log15.InitWithConfig(&logger.LogConfig{
//		Console: logger.ConsoleLogConfig{
//			Level:    4,
//			Modules:  "*",
//			ShowPath: false,
//			UseColor: true,
//			Console:  true,
//		},
//		File: logger.FileLogConfig{
//			Level:    4,
//			Modules:  "*",
//			Save:     true,
//			FilePath: "./logs",
//			FileName: "logger.json",
//		},
//	})
//
//	startTime := logger.CurrentTime()
//
//	log := logger.New("log15")
//
//	for i := 0; i < 2; i++ {
//		log.Info("=========================", "i", i)
//		for j := 0; j < 10; j++ {
//			log.Debug("test1 debug", "i", i, "j", j)
//			if j%3 == 0 {
//				log.Info("test2 info", "i", i, "j", j)
//			}
//			if j%5 == 0 {
//				logger.Error("test2 info", "i", i, "j", j)
//			}
//		}
//	}
//	logger.Info("总耗时", "elapsed", logger.CurrentTime()-startTime)
//}

func zapLog() {
	zap.InitWithConfig(&logger.LogConfig{
		Console: logger.ConsoleLogConfig{
			Level:    4,
			Modules:  "*",
			ShowPath: false,
			UseColor: true,
			Console:  true,
		},
		File: logger.FileLogConfig{
			Level:    4,
			Modules:  "*",
			Save:     true,
			FilePath: "./logs",
			FileName: "logger.json",
		},
	})

	startTime := logger.CurrentTime()

	log := logger.New("zap")

	for i := 0; i < 2; i++ {
		log.Info("=========================", "i", i)
		for j := 0; j < 10; j++ {
			log.Debug("test1 debug", "i", i, "j", j)
			if j%3 == 0 {
				log.Info("test2 info", "i", i, "j", j)
			}
			if j%5 == 0 {
				logger.Error("test2 info", "i", i, "j", j)
			}
		}
	}
	logger.Info("总耗时", "elapsed", logger.CurrentTime()-startTime)
}

func logrusLog() {
	logrus.InitWithConfig(&logger.LogConfig{
		Console: logger.ConsoleLogConfig{
			Level:    4,
			Modules:  "*",
			ShowPath: false,
			UseColor: true,
			Console:  true,
		},
		File: logger.FileLogConfig{
			Level:    4,
			Modules:  "*",
			Save:     true,
			FilePath: "./logs",
			FileName: "logger.json",
		},
	})

	startTime := logger.CurrentTime()

	log := logger.New("logrus")

	for i := 0; i < 2; i++ {
		log.Info("=========================", "i", i)
		for j := 0; j < 10; j++ {
			log.Debug("test1 debug", "i", i, "j", j)
			if j%3 == 0 {
				log.Info("test2 info", "i", i, "j", j)
			}
			if j%5 == 0 {
				logger.Error("test2 info", "i", i, "j", j)
			}
		}
	}
	logger.Info("总耗时", "elapsed", logger.CurrentTime()-startTime)
}
