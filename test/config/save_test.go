package config

import (
    "gourd/internal/config"
    "testing"
)

// 保存配置文件测试
func TestSaveConfig(t *testing.T) {

    err := config.SetConfigPath("./configs")
    if err != nil {
        t.Errorf("set config path error: %v", err)
        return
    }

    var conf = map[string]config.DbConfig{
        "mysql": {
            Database:    "test",
            Host:        "127.0.0.1",
            Param:       "charset=utf8mb4&parseTime=True&loc=Local",
            Pass:        "123456",
            Port:        3306,
            SlowLogTime: 1000,
            Type:        "mysql",
            User:        "root",
        },
    }

    err = config.Marshal("database", conf)
    if err != nil {
        t.Errorf("save config error: %v", err)
        return
    }
}
