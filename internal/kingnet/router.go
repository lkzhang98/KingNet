package kingnet

import (
	"go.mod/internal/kingnet/iface"
	"go.mod/internal/kingnet/knet"
	"go.mod/internal/pkg/log"
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
