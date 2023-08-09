// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package kingnet

import (
	"KingNet/internal/kingnet/iface"
	"KingNet/internal/kingnet/knet"
	"KingNet/internal/pkg/log"
)

type PingRouter struct {
	knet.BaseRouter
}

func (this *PingRouter) Handle(request iface.RequestI) {
	log.Debug("Call PingRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	log.Debugf("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		log.Error(err)
	}
}

type HelloKingNetRouter struct {
	knet.BaseRouter
}

func (h *HelloKingNetRouter) Handle(request iface.RequestI) {
	log.Debug("Call HelloKingNetRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	log.Debugf("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(1, []byte("Hello KingNet Router V0.8"))
	if err != nil {
		log.Error(err)
	}
}
