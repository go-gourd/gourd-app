package cmd

import (
	"fmt"
	"github.com/go-gourd/gourd/cmd"
)

// Register 注册命令入口
func Register() {

	//默认命令行操作
	cmd.SetDefault(cmd.Commend{
		Handler: func(args []string) {
			//这里直接调用内置 start 命令
			_ = cmd.Exec("start", []string{
				args[0],
				"start",
			})
		},
	})

	//命令行示例
	cmd.Add(cmd.Commend{
		Name:    "test",
		Explain: "This is a test template.",
		Handler: func(args []string) {
			fmt.Println("Test command run successfully.")
		},
	})

}
