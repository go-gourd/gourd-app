package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 设置日志级别为Debug
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	log.Print("This is a standard log message 1")

	log.Print("This is a standard log message 2")

	log.Print("This is a standard log message 3")

	slog.Debug("This is a debug message")
	slog.Info("This is a info message")
	slog.Warn("This is a warn message")
	slog.Error("This is a error message")

}
