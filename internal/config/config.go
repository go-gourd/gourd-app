package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

// 默认配置文件目录
const defaultPath = "./configs"

// 配置文件目录
var configDir = defaultPath

// SetConfigPath 设置文件目录
func SetConfigPath(path string) {
	configDir = path
}

// Unmarshal 读取自定义配置文件
func Unmarshal(name string, v any) error {

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
