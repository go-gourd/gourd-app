package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/ghttp"
	"gourd/app/dal/query"
)

type HandlerUser struct {
}

func (*HandlerUser) GetUser(c *gin.Context) {

	first, _ := query.User.Where(query.User.ID.Neq(99)).Find()
	fmt.Println(first)

	ghttp.Success(c, "test", first)
}

func Info(c *gin.Context) {

	data := make(map[string]any)
	data["list"] = "666"

	ghttp.Success(c, "test", data)
}
