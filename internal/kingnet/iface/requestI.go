// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

// RequestI defines the method to process requests.
type RequestI interface {
	// GetConnection returns the connection
	GetConnection() ConnectionI
	// GetData returns the data
	GetData() []byte
	// GetMsgID returns the message ID
	GetMsgID() uint32
}
