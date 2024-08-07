package common

import (
	"encoding/json"
	"net/http"
)

// BaseCtl 基础控制器
// 所有控制器都应继承此控制器，可以在此控制器中定义公共方法
type BaseCtl struct {
}

// Success 成功时响应
func (*BaseCtl) Success(w http.ResponseWriter, message string, data any) (err error) {
	if message == "" {
		message = "success"
	}
	res := map[string]any{
		"code":    0,
		"data":    data,
		"message": message,
	}

	str, _ := json.Marshal(res)
	_, err = w.Write(str)
	return
}

// Fail 失败响应
func (*BaseCtl) Fail(w http.ResponseWriter, code int, message string, data any) (err error) {
	if message == "" {
		message = "fail"
	}
	res := map[string]any{
		"code":    code,
		"data":    data,
		"message": message,
	}

	str, _ := json.Marshal(res)
	_, err = w.Write(str)
	return
}
