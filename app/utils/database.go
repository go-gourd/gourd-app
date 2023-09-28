package utils

import (
	"errors"
	"github.com/go-gourd/database"
	"github.com/go-gourd/gourd/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gourd/app/orm/query"
	"time"
)

type gormLogWriter struct{}

func (w gormLogWriter) Printf(format string, args ...any) {
	log.Warnf(format, args...)
}

// InitDatabase 初始化数据库连接
func InitDatabase() (err error) {

	// 连接数据库
	dbConfig := database.GetConfig("mysql")
	if dbConfig == nil {
		return errors.New("mysql config is nil")
	}

	// 接管gorm日志
	newLogger := &gorm.Config{
		Logger: logger.New(
			gormLogWriter{},
			logger.Config{
				SlowThreshold:             time.Duration(dbConfig.SlowLogTime) * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Error,                                           // 日志级别
				IgnoreRecordNotFoundError: true,                                                   // 忽略记录未找到错误
				Colorful:                  false,                                                  // 禁用彩色打印
			},
		),
	}

	mysqlDb, err := gorm.Open(mysql.Open(dbConfig.GenerateDsn()), newLogger)
	if err != nil {
		return err
	}
	query.SetDefault(mysqlDb)

	return nil
}
