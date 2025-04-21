package config

import (
    "gourd/internal/config"
    "testing"
)

func TestReadConfig(t *testing.T) {
    dbConfig := &config.DbConfigMap{}

    err := config.SetConfigPath("./configs")
    if err != nil {
        t.Errorf("set config path error: %v", err)
        return
    }

    err = config.Unmarshal("database", dbConfig)
    if err != nil {
        t.Errorf("Unmarshal error: %v", err)
        return
    }

    t.Log(dbConfig)

    // 读取指定数据库配置
    mysqlDb, err := config.GetDBConfig("mysql")
    if err != nil {
        return
    }

    // 设置指定数据库配置
    err = config.SetDBConfig("mysql", mysqlDb)
    if err != nil {
        return
    }

}
