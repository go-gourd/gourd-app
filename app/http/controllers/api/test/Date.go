package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd"
	"time"
)

type Date struct {
	gourd.BaseController
}

func (con Date) Get(c *gin.Context) {
	con.Success(c, time.Now().String())
}
