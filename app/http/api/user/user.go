package user

import (
	"app/service/utils"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {

	data := make(map[string]any)
	data["list"] = "666"

	utils.ResJson(c, 0, "test", data)
}
