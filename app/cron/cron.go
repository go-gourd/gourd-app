package cron

import (
	"fmt"
	"github.com/go-gourd/gourd/cron"
)

func RegisterCron() {

	_ = cron.Add("* * * * *", func() {
		fmt.Println("定时任务测试")
	})

}