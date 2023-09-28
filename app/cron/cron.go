package cron

import (
	"fmt"
	"github.com/go-gourd/gourd/cron"
	"time"
)

// RegisterCron 定时任务注册
func RegisterCron() {

	_ = cron.Add("* * * * *", func() {
		fmt.Printf("定时任务示例 %s\n", time.Now().Format(time.TimeOnly))
	})

}
