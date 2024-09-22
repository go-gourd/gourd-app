package initialize

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gourd/internal/config"
	"gourd/internal/orm/query"
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
		return errors.New("database.mysql config is nil")
	}

	// 连接数据库
	dsn := dbConf.GenerateDsn()
	gormConfig := &gorm.Config{
		// 替换默认日志
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
	mysqlDb, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return err
	}

	// 设置默认查询器
	query.SetDefault(mysqlDb)

	return nil
}
