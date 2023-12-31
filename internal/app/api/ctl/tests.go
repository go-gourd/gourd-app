package controller

import (
	"github.com/go-gourd/gourd/event"
	"gourd/internal/pkg/libhttp"
	"net/http"
)

// TestsController 测试
type TestsController struct {
	libhttp.BaseController //继承基础控制器
}

// Test 测试
func (ct *TestsController) Test(w http.ResponseWriter, _ *http.Request) {

	event.Trigger("test", "test value")

	// 响应结果
	_ = ct.Success(w, "", nil)
}
