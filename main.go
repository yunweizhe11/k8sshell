package main

import (
	"fmt"

	"github.com/kevin/chapter4/k8sshell/src/cmds"
	// "./cmds"
)

func main() {
	// 获取根命令
	rootCmd := cmds.GetRootCommand()

	// 执行命令
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
}
