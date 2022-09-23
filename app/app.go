// Package app
// @Title  应用管理
package app

import (
	"ginapp/app/http/middleware"
	"ginapp/app/http/router"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml/v2"
)

// StartApp 启动应用
func StartApp() {

	GetConfig()

	app := gin.Default()

	//注册HTTP中间件
	middleware.RegisterMiddleware(app)

	//注册HTTP路由
	router.RegisterRoute(app)

	err := app.Run()
	if err != nil {
		return
	}

}

func GetConfig() {
	type MyConfig struct {
		Version int
		Name    string
		Tags    []string
	}
	doc := `
version = 2
name = "go-toml"
tags = ["go", "toml"]
`

	var cfg MyConfig
	err := toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		panic(err)
	}

}
