package middleware

import (
	"net/http"
	"os"
	"path/filepath"
)

// StaticOrNext 静态文件处理
func StaticOrNext(dir string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fPath := filepath.Join(dir, r.URL.Path)

			info, err := os.Stat(fPath)
			if err == nil {
				if info.IsDir() {
					// 如果是目录，判断是否存在 index.html
					indexPath := filepath.Join(fPath, "index.html")
					if _, err := os.Stat(indexPath); err == nil {
						http.ServeFile(w, r, indexPath)
						return
					}
					next.ServeHTTP(w, r)
					return
				}
				// 如果是文件，直接返回，避免多余的 301
				http.ServeFile(w, r, fPath)
				return
			}

			// 文件不存在，继续走 chi 路由
			next.ServeHTTP(w, r)
		})
	}
}
