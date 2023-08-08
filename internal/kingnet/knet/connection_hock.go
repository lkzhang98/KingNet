package knet

import (
	"go.mod/internal/kingnet/iface"
	"go.mod/internal/pkg/log"
)

// DoConnectionBegin 连接开始的时候执行
func DoConnectionBegin(conn iface.ConnectionI) {
	log.Info("DoConnectionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		log.Error(err)
	}
}

// DoConnectionLost 连接断开的时候执行
func DoConnectionLost(conn iface.ConnectionI) {
	log.Info("DoConnectionLost is Called ... ")
}
