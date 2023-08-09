// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import (
	"KingNet/internal/kingnet/iface"
	"KingNet/internal/pkg/log"
)

// DoConnectionBegin 连接开始的时候执行.
func DoConnectionBegin(conn iface.ConnectionI) {
	log.Info("DoConnectionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		log.Error(err)
	}
}

// DoConnectionLost 连接断开的时候执行.
func DoConnectionLost(conn iface.ConnectionI) {
	log.Info("DoConnectionLost is Called ... ")
}
