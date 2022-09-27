package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionMiddle 获取SESSION中间件
func SessionMiddle() gin.HandlerFunc {

	// 创建基于 cookie 的存储引擎，该参数是用于加密的密钥
	store := cookie.NewStore([]byte("gourd"))

	store.Options(sessions.Options{
		Path:   "./runtime/",
		MaxAge: 86400, //过期时间，秒
	})

	// 设置Session中间件，参数1是SessionName，参数2是储存引擎
	sess := sessions.Sessions("session", store)

	return sess
}
