package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/ghttp"
)

type HandlerUser struct {
}

func (*HandlerUser) GetUser(c *gin.Context) {
	data := make(map[string]any)
	data["info"] = "test data"

	ghttp.Success(c, 0, "test", data)
}

func Info(c *gin.Context) {

	data := make(map[string]any)
	data["list"] = "666"

	ghttp.Success(c, 0, "test", data)
}
