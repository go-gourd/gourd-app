package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

// 默认配置文件目录
const defaultPath = "./configs"

// 配置文件目录
var configDir = ""

// SetConfigPath 设置文件目录
func SetConfigPath(path string) error {
	configDir = path
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return os.MkdirAll(configDir, os.ModePerm)
	}
	return nil
}

// Exists 判断配置是否存在
func Exists(name string) bool {

	if configDir == "" {
		err := SetConfigPath(defaultPath)
		if err != nil {
			return false
		}
	}

	var file = configDir + "/" + name + ".toml"
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// Unmarshal 读取自定义配置文件
func Unmarshal(name string, v any) error {

	if configDir == "" {
		return SetConfigPath(defaultPath)
	}

	var file = configDir + "/" + name + ".toml"
	tomlData, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(tomlData, v)
	if err != nil {
		return err
	}
	return nil
}

// Marshal 写入自定义配置文件
func Marshal(name string, v any) error {

	if configDir == "" {
		return SetConfigPath(defaultPath)
	}

	var file = configDir + "/" + name + ".toml"
	tomlData, err := toml.Marshal(v)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, tomlData, 0644)
	if err != nil {
		return err
	}
	return nil
}
