package main

import (
	"ginlearn/config"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(cfg *config.Config) {

	//创建一个新的日志记录器
	Log = logrus.New()

	//设置日志级别
	level := logrus.InfoLevel

	if cfg != nil && cfg.LogLevel != "" {

		//解析日志级别
		parsedLevel, err := logrus.ParseLevel(cfg.LogLevel)

		if err != nil {
			Log.Warnf("无效的日志级别 '%s'，使用默认级别 'info'", cfg.LogLevel)
		} else {
			level = parsedLevel
		}
	}

	//设置日志级别为info
	Log.SetLevel(level)

	/// 设置为文本格式输出（默认）
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,                  // 显示完整时间
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
		ForceColors:     true,                  // 强制彩色输出
	})

	// 设置标准输出
	Log.SetOutput(os.Stdout)

	Log.Info("日志系统初始化完成")
}
