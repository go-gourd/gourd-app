package base

import (
	"encoding/json"
	"net/http"
)

// BaseController 基础控制器
type BaseController struct {
}

// Success 成功时响应
func (*BaseController) Success(w http.ResponseWriter, message string, data any) (err error) {
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
func (*BaseController) Fail(w http.ResponseWriter, code int, message string, data any) (err error) {
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