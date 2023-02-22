package cmd

import (
	"fmt"
	"github.com/go-gourd/gourd/cmd"
)

func RegisterCmd() {

	//命令行示例
	cmd.AddCmd(cmd.Commend{
		Name:    "test",
		Explain: "This is a test template.",
		Handler: func(args []string) {
			fmt.Println("Test command run successfully.")
		},
	})

}
