package controller

import (
	"github.com/go-gourd/gourd/event"
	"gourd/internal/http/api/common"
	"net/http"
)

// TestsController 测试
type TestsController struct {
	common.BaseController //继承基础控制器
}

// Test 测试
func (ctl *TestsController) Test(w http.ResponseWriter, _ *http.Request) {

	event.Trigger("test", nil)

	// 响应结果
	_ = ctl.Success(w, "", nil)
}
