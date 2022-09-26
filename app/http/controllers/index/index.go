package index

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd"
)

type Controller struct {
	gourd.BaseController
}

func (con Controller) Index(c *gin.Context) {
	con.Success(c, "index.Index")
}
