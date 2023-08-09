// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

// DataPackI 定义拆包和封包方法.
type DataPackI interface {
	GetHeadLen() uint32                // 获取包头长度方法
	Pack(msg MessageI) ([]byte, error) // 封包方法
	Unpack([]byte) (MessageI, error)   // 拆包方法
}
