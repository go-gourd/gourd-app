package initialize

import (
	"app/internal/config"
	"app/internal/global"
	"app/internal/orm/query"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

// 使用自定义logger接管gorm日志
type dbLogWriter struct{}

func (w dbLogWriter) Printf(format string, args ...any) {
	slog.Warn(fmt.Sprintf(format, args...))
}

// InitDatabase 初始化数据库连接
func InitDatabase() error {

	// 连接数据库
	dbConf, err := config.GetDBConfig("mysql")
	if err != nil {
		return errors.New("mysql config is nil")
	}

	// 连接数据库
	gormConfig := &gorm.Config{
		Logger: logger.New(
			dbLogWriter{},
			logger.Config{
				SlowThreshold:             time.Duration(dbConf.SlowLogTime) * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Warn,                                          // 日志级别
				IgnoreRecordNotFoundError: true,                                                 // 忽略记录未找到错误
				Colorful:                  true,                                                 // 禁用彩色打印
			},
		),
	}
	mysqlDb, err := gorm.Open(mysql.Open(dbConf.GenerateDsn()), gormConfig)
	if err != nil {
		return err
	}

	// 设置全局数据库连接
	global.SetDb("mysql", mysqlDb)

	// 设置默认查询器
	query.SetDefault(mysqlDb)

	return nil
}
