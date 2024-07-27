package util

import (
	"context"
	"gopkg.in/natefinch/lumberjack.v2"
	"gourd/internal/config"
	"io"
	"log/slog"
	"os"
	"time"
)

// CustomHandler 自定义日志处理器
type CustomHandler struct {
	Level  slog.Level
	Writer io.Writer
}

func (CustomHandler) WithAttrs([]slog.Attr) slog.Handler { return nil }

func (CustomHandler) WithGroup(string) slog.Handler { return nil }

func (h CustomHandler) Enabled(_ context.Context, l slog.Level) bool {
	return l >= h.Level
}

func (h CustomHandler) Handle(_ context.Context, r slog.Record) error {

	dt := time.Now().Format("2006-01-02 15:04:05")
	msg := dt + " " + r.Level.String() + " " + r.Message

	// 输出日志属性
	r.Attrs(func(a slog.Attr) bool {
		msg += " " + a.Key + "=" + a.Value.String()
		return true
	})
	_, err := h.Writer.Write([]byte(msg + "\n"))
	return err
}

// InitLog 初始化日志
func InitLog() {

	conf, err := config.GetLogConfig()
	if err != nil {
		panic(err)
	}

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
	err = level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		level.Set(slog.LevelInfo) // 默认日志级别为info
	}

	options := &slog.HandlerOptions{
		Level: level.Level(), // 设置日志级别
	}

	var handler slog.Handler
	if conf.Encoding == "json" {
		handler = slog.NewJSONHandler(logWriter, options)
	} else if conf.Encoding == "text" {
		handler = slog.NewTextHandler(logWriter, options)
	} else if conf.Encoding == "default" {
		handler = CustomHandler{
			Level:  level.Level(),
			Writer: logWriter,
		}
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
