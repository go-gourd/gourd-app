package config

import (
	"errors"
	"fmt"
)

// DbConfig 适用于单个连接的配置
type DbConfig struct {
	Type        string `toml:"type" json:"type" comment:"数据库类型" `
	Host        string `toml:"host" json:"host" comment:"数据库地址"`
	Port        int    `toml:"port" json:"port" comment:"端口"`
	User        string `toml:"user" json:"user" comment:"用户"`
	Pass        string `toml:"pass" json:"pass" comment:"密码"`
	Database    string `toml:"database" json:"database" comment:"数据库名"`
	Param       string `toml:"param" json:"param" comment:"连接参数"`
	SlowLogTime int    `toml:"slow_log_time" json:"slow_log_time" comment:"慢日志阈值（毫秒）0为不开启"`
}

// DbConfigMap 适用于多个连接的配置
type DbConfigMap map[string]DbConfig

// GenerateDsn 根据配置生成sdn连接信息
func (conf DbConfig) GenerateDsn() string {
	dsn := ""
	dsnParam := ""

	if conf.Type == "mysql" {
		if conf.Param != "" {
			dsnParam = "?" + conf.Param
		}
		dsnF := "%s:%s@(%s:%d)/%s%s"
		dsn = fmt.Sprintf(dsnF, conf.User, conf.Pass, conf.Host, conf.Port, conf.Database, dsnParam)
	} else if conf.Type == "sqlserver" {
		if conf.Param != "" {
			dsnParam = "&" + conf.Param
		}
		dsnF := "sqlserver://%s:%s@%s:%d?database=%s%s"
		dsn = fmt.Sprintf(dsnF, conf.User, conf.Pass, conf.Host, conf.Port, conf.Database, dsnParam)
	} else if conf.Type == "postgres" {
		dsnF := "host=%s user=%s password=%s dbname=%s port=%d %s " + conf.Param
		dsn = fmt.Sprintf(dsnF, conf.Host, conf.User, conf.Pass, conf.Database, conf.Port, dsnParam)
	} else if conf.Type == "oracle" {
		dsnF := "%s/%s@%s:%d/%s"
		dsn = fmt.Sprintf(dsnF, conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	}

	return dsn
}

var dbConfig *DbConfigMap

// GetDBConfigAll 获取所有数据库配置
func GetDBConfigAll() (*DbConfigMap, error) {
	key := "database"

	if dbConfig == nil {
		dbConfig = &DbConfigMap{}
	}

	// 如果配置不存在，则创建默认配置
	if !Exists(key) {
		err := SetDBConfigAll(dbConfig)
		if err != nil {
			return nil, err
		}
	}

	err := Unmarshal(key, dbConfig)
	if err != nil {
		return nil, err
	}
	return dbConfig, nil
}

// SetDBConfigAll 设置所有数据库配置
func SetDBConfigAll(conf *DbConfigMap) error {
	key := "database"
	dbConfig = conf
	return Marshal(key, conf)
}

// GetDBConfig 获取指定数据库配置
func GetDBConfig(name string) (*DbConfig, error) {

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

// SetDBConfig 设置指定数据库配置
func SetDBConfig(name string, conf *DbConfig) error {

	all, err := GetDBConfigAll()
	if err != nil {
		return err
	}

	allDb := *all

	// 判断all中是否存在
	if _, ok := allDb[name]; ok {
		allDb[name] = *conf
	} else {
		allDb[name] = *conf
	}

	return SetDBConfigAll(&allDb)
}
