package index

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd"
	"github.com/go-gourd/gourd/log"
)

type Controller struct {
	gourd.BaseController
}

func (con Controller) Index(c *gin.Context) {
	con.Success(c, "index.Index")
}

func (con Controller) Session(c *gin.Context) {

	// 初始化 session 对象
	session := sessions.Default(c)

	data := session.Get("data")
	if data != "world" {
		session.Set("data", "world")
		err := session.Save()
		if err != nil {
			log.Error(err.Error())
		}
	}

	if data != nil {
		log.Info(data.(string))
	}
}
