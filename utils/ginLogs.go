package utils

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"runtime"
	"strings"
)

func InitLogger() error {

	logFile := &lumberjack.Logger{
		Filename:   "E:\\goPro\\goPro4\\logs\\logs.log",
		MaxSize:    2,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	// 设置日志输出
	log.SetOutput(logFile)

	// 设置日志格式
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	return nil
}

// Info 记录信息级别日志
func Info(v ...interface{}) {
	logWithCaller("INFO", v...)
}

// Error 记录错误级别日志
func Error(v ...interface{}) {
	logWithCaller("ERROR", v...)
}

func logWithCaller(level string, v ...interface{}) {
	// 获取调用者的文件名和行号
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Println("Failed to get caller information")
		return
	}

	// 提取文件名
	fileName := strings.TrimPrefix(file, "/")

	// 构造日志消息
	msg := fmt.Sprintf("%s %s:%d: %v", level, fileName, line, v)

	// 输出日志
	log.Println(msg)
}

//func InitLogger01() {
//	logger := logrus.New()
//	//logger.SetFormatter(&logrus.TextFormatter{
//	//	FullTimestamp: true,
//	//})
//
//	logger.SetLevel(logrus.ErrorLevel)
//	//	需要安装这个模块
//	//go get gopkg.in/natefinch/lumberjack.v2
//	hook := lumberjack.Logger{
//		Filename:   "E:\\goPro\\goPro4\\logs\\logs.log",
//		MaxSize:    1,
//		MaxBackups: 2,
//		MaxAge:     28,
//	}
//	logger.Out = &hook
//	logger.WithFields(logrus.Fields{
//		"animal": "walrus",
//	}).Infof("User logged in")
//
//	logger.WithFields(logrus.Fields{
//		"key": "value",
//	}).WithError(errors.New("报错了!!!!"))
//}
