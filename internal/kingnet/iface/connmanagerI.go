// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

type ConnManagerI interface {
	// Add adds a connection to the connection manager.
	Add(conn ConnectionI)
	// Remove removes a connection from the connection manager.
	Remove(conn ConnectionI)
	// Get gets a connection from the connection manager.
	Get(connID uint32) (ConnectionI, error)
	// Len returns the number of connections in the connection manager.
	Len() int
	// ClearConn removes all connections from the connection manager.
	ClearConn()
}
