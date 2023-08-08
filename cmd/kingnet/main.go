package main

import (
	"go.mod/internal/kingnet"
	"os"
)

// main 是整个KingNet的主程序入口。
func main() {
	command := kingnet.NewKingNetCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
