package config

import (
    "errors"
    "github.com/go-gourd/database"
)

var dbConfig *database.ConfigMap

// GetDBConfigAll 获取所有数据库配置
func GetDBConfigAll() (*database.ConfigMap, error) {

    if dbConfig == nil {
        dbConfig = &database.ConfigMap{}
    }

    err := Unmarshal("database", dbConfig)
    if err != nil {
        return nil, err
    }
    return dbConfig, nil
}

// GetDBConfig 获取指定数据库配置
func GetDBConfig(name string) (*database.Config, error) {

    all, err := GetDBConfigAll()
    if err != nil {
        return nil, err
    }

    allDb := *all

    // 判断all中是否存在
    if _, ok := allDb[name]; ok {
        db := allDb[name]
        return &db, nil
    }
    return nil, errors.New("database config not found")
}
