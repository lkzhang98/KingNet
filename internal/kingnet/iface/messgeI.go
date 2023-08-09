// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

type MessageI interface {
	GetDataLen() uint32 // 获取消息数据段长度
	GetMsgId() uint32   // 获取消息ID
	GetData() []byte    // 获取消息内容

	SetMsgId(uint32)   // 设计消息ID
	SetData([]byte)    // 设计消息内容
	SetDataLen(uint32) // 设置消息数据段长度
}
