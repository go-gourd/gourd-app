package controller

import (
	"app/internal/http/api/controller/common"
	"github.com/go-gourd/gourd/event"
	"net/http"
)

// TestsCtl 测试
type TestsCtl struct {
	common.Base //继承基础控制器
}

// Test 测试
func (ctl *TestsCtl) Test(w http.ResponseWriter, r *http.Request) {

	event.Trigger("test", r.Context())

	// 响应结果
	_ = ctl.Success(w, "", nil)
}
