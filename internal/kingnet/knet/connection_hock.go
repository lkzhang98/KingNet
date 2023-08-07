package knet

import (
	"go.mod/internal/kingnet/iface"
	"go.mod/internal/pkg/log"
)

func DoConnectionBegin(conn iface.ConnectionI) {
	log.Info("DoConnecionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		log.Error(err)
	}
}

// 连接断开的时候执行
func DoConnectionLost(conn iface.ConnectionI) {
	log.Info("DoConneciotnLost is Called ... ")
}
