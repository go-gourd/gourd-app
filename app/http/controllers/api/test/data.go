package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd"
	"github.com/go-gourd/gourd/database"
	"gourd/app/model"
)

type Data struct {
	gourd.BaseController
}

func (con Data) GetUser(c *gin.Context) {

	db := database.GetDb()

	var user model.User

	db.First(&user)

	con.Success(c, user.Nickname)

}
