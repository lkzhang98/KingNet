// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import "KingNet/internal/kingnet/iface"

type Request struct {
	conn iface.ConnectionI // 已经和客户端建立好的 链接
	msg  iface.MessageI
}

var _ iface.RequestI = &Request{}

func (r *Request) GetConnection() iface.ConnectionI {
	return r.conn
}

// GetData 获取请求消息的数据.
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求的消息的ID.
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
