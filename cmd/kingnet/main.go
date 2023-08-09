// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package main

import (
	"os"

	"KingNet/internal/kingnet"
)

// main 是整个KingNet的主程序入口。
func main() {
	command := kingnet.NewKingNetCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
