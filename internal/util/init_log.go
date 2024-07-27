package util

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"gourd/internal/config"
	"io"
	"log/slog"
	"os"
)

func InitLog() {

	conf, _ := config.GetLogConfig()

	// 日志文件输出
	output := &lumberjack.Logger{
		Filename:  conf.LogFile,
		MaxSize:   conf.MaxSize, // megabytes
		LocalTime: true,
	}

	writers := []io.Writer{output}
	if conf.Console {
		// 日志是否输出到控制台
		writers = append(writers, os.Stdout)
	}
	logWriter := io.MultiWriter(writers...)

	// 设置日志级别
	level := slog.LevelVar{}
	err := level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		level.Set(slog.LevelInfo) // 默认日志级别为info
	}

	options := &slog.HandlerOptions{
		Level: level.Level(), // 设置日志级别
	}

	var handler slog.Handler
	if conf.Encoding == "json" {
		handler = slog.NewJSONHandler(logWriter, options)
	} else {
		handler = slog.NewTextHandler(logWriter, options)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)

}
