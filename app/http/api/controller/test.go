package controller

import (
	"github.com/go-gourd/gourd/event"
	"gourd/app/http/base"
	"net/http"
)

// TestController 测试
type TestController struct {
	base.BaseController //继承基础控制器
}

// Test 测试
func (ct *TestController) Test(w http.ResponseWriter, _ *http.Request) {

	event.Trigger("test", "test value")

	// 响应结果
	_ = ct.Success(w, "", nil)
}
