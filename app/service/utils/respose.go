package utils

import "github.com/gin-gonic/gin"

func ResJson(c *gin.Context, code int, message string, data any) {
	res := gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	}
	c.JSON(200, res)
}
