package cron

import (
    "github.com/go-gourd/gourd/cron"
    "log/slog"
    "time"
)

// Register 定时任务注册
func Register() {

    _ = cron.Add("* * * * *", func() {
        slog.Info("定时任务示例" + time.Now().Format(time.TimeOnly))
    })

}
