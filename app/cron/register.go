package cron

import (
	"github.com/go-gourd/gourd/cron"
	"github.com/go-gourd/gourd/log"
)

// RegisterCron 注册定时任务
func RegisterCron() {

	//注册定时任务
	err := cron.AddTask("* * * * *", func() {
		log.Debug("TestCron.")
	})
	if err != nil {
		log.Error(err.Error())
	}

}
