package tools

import (
    "errors"
    "fmt"
    "github.com/go-gourd/database"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gourd/internal/config"
    "gourd/internal/orm/query"
    "log/slog"
    "time"
)

// 使用自定义logger接管gorm日志
type customLogWriter struct{}

func (w customLogWriter) Printf(format string, args ...any) {
    slog.Warn(fmt.Sprintf(format, args...))
}

// 初始化logger
func initDbLogger(dbConfig *database.Config) logger.Interface {
    return logger.New(
        customLogWriter{},
        logger.Config{
            SlowThreshold:             time.Duration(dbConfig.SlowLogTime) * time.Millisecond, // 慢 SQL 阈值
            LogLevel:                  logger.Error,                                           // 日志级别
            IgnoreRecordNotFoundError: true,                                                   // 忽略记录未找到错误
            Colorful:                  false,                                                  // 禁用彩色打印
        },
    )
}

// InitDatabase 初始化数据库连接
func InitDatabase() error {

    // 连接数据库
    dbConf, err := config.GetDBConfig("mysql")
    if err != nil {
        return errors.New("mysql config is nil")
    }

    // 连接数据库
    dsn := dbConf.GenerateDsn()
    gormConfig := &gorm.Config{
        Logger: initDbLogger(dbConf),
    }
    mysqlDb, err := gorm.Open(mysql.Open(dsn), gormConfig)
    if err != nil {
        return err
    }

    // 设置默认查询器
    query.SetDefault(mysqlDb)

    return nil
}
