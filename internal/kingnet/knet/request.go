package knet

import "go.mod/internal/kingnet/iface"

type Request struct {
	conn iface.ConnectionI //已经和客户端建立好的 链接
	msg  iface.MessageI
}

var _ iface.RequestI = &Request{}

func (r *Request) GetConnection() iface.ConnectionI {
	return r.conn
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
