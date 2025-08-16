package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {

	// 创建新的日志实例
	Log = logrus.New()

	/**
	logrus.InfoLevel 是 Logrus 日志库中定义的一个日志级别常量，
	表示 "信息" 级别，用于标记常规的运行时信息日志。
	*/
	level := logrus.InfoLevel

	// 设置全局日志级别为 InfoLevel
	Log.SetLevel(level)

	// 设置为文本格式输出（默认）
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,                  // 显示完整时间
		TimestampFormat: "2006-01-02 15:04:05", // 自定义时间格式
		ForceColors:     true,                  // 强制彩色输出
	})

	// 设置标准输出
	Log.SetOutput(os.Stdout)

	Log.Info("日志系统初始化完成")
}
