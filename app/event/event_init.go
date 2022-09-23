package event

import (
	"fmt"
	"ginapp/app/http/router"
	"github.com/go-gourd/gourd"
)

// InitEvent 初始化时执行
func InitEvent() {

	app := gourd.GetGinEngine()

	fmt.Println("InitEvent")
	fmt.Println("InitEvent")

	//注册路由
	router.RegisterRoute(app)

}
