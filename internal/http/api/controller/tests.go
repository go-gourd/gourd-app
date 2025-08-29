package controller

import (
	"net/http"

	"github.com/go-gourd/gourd/event"
)

// TestsCtl 测试
type TestsCtl struct {
	Base //继承基础控制器
}

// Test 测试
func (ctl *TestsCtl) Test(w http.ResponseWriter, r *http.Request) {

	event.Trigger("test", r.Context())

	// 响应结果
	_ = ctl.Success(w, "", nil)
}
