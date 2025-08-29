package initialize

import (
	"app/internal/config"
	"app/internal/orm/query"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 使用自定义logger接管gorm日志
type dbLogWriter struct{}

func (w dbLogWriter) Printf(format string, args ...any) {
	slog.Warn(fmt.Sprintf(format, args...))
}

// InitDatabase 初始化数据库连接
func InitDatabase() error {

	// 连接数据库
	dbConf, err := config.GetDBConfig("main")
	if err != nil {
		return errors.New("database.main config is nil")
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

	var mainDb *gorm.DB

	if dbConf.Type == "mysql" {
		mainDb, err = gorm.Open(mysql.Open(dbConf.GenerateDsn()), gormConfig)
		if err != nil {
			return err
		}
	} else {
		return errors.New("database type is not supported")
	}

	// 设置默认查询器
	query.SetDefault(mainDb)

	return nil
}
