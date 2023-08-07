package main

import (
	"go.mod/internal/kingnet"
	"os"
)

func main() {
	command := kingnet.NewKingNetCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
