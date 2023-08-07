package iface

type MsgHandleI interface {
	// DoMsgHandler handles a message
	DoMsgHandler(request RequestI)
	// AddRouter add router
	AddRouter(msgId uint32, router RouterI)
	// StartWorkerPool starts a worker pool
	StartWorkerPool()
	// SendMsgToTaskQueue sends a message to task queue
	SendMsgToTaskQueue(request RequestI)
}
