// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

type MsgHandlerI interface {
	// DoMsgHandler handles a message
	DoMsgHandler(request RequestI)
	// AddRouter add router
	AddRouter(msgId uint32, router RouterI)
	// StartWorkerPool starts a worker pool
	StartWorkerPool()
	// SendMsgToTaskQueue sends a message to task queue
	SendMsgToTaskQueue(request RequestI)
}
