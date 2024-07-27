package controller

import (
	"github.com/go-gourd/gourd/event"
	"gourd/internal/http/api/common"
	"net/http"
)

// TestsCtl 测试
type TestsCtl struct {
	common.BaseCtl //继承基础控制器
}

// Test 测试
func (ctl *TestsCtl) Test(w http.ResponseWriter, r *http.Request) {

	event.Trigger("test", r.Context())

	// 响应结果
	_ = ctl.Success(w, "", nil)
}
