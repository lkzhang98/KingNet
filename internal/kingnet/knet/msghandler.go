// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import (
	"strconv"

	"KingNet/internal/kingnet/iface"
	"KingNet/internal/pkg/log"
)

// MsgHandler 定义了消息的处理方法.
type MsgHandler struct {
	Apis           map[uint32]iface.RouterI
	WorkerPoolSize uint32
	TaskQueue      []chan iface.RequestI // 存放每个MsgId 所对应的处理方法的map属性
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]iface.RouterI),
		WorkerPoolSize: ServerOption.WorkerPoolSize,
		TaskQueue:      make([]chan iface.RequestI, ServerOption.MaxWorkerTaskLen),
	}
}

var _ iface.MsgHandlerI = &MsgHandler{}

// DoMsgHandler 马上以非阻塞方式处理消息.
func (mh *MsgHandler) DoMsgHandler(request iface.RequestI) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		log.Error("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	// 执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体的处理逻辑.
func (mh *MsgHandler) AddRouter(msgId uint32, router iface.RouterI) {
	// 1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	// 2 添加msg与api的绑定关系
	mh.Apis[msgId] = router
	log.Debug("Add api msgId = ", msgId)
}

// StartOneWorker 开启一个worker工作.
func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan iface.RequestI) {
	log.Info("Worker ID = ", workerID, " is started.")
	// 不断的等待队列中的消息
	for {
		select {
		// 有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// StartWorkerPool 启动worker工作池.
func (mh *MsgHandler) StartWorkerPool() {
	// 遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan iface.RequestI, ServerOption.MaxWorkerTaskLen)
		// 启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// SendMsgToTaskQueue 向任务队列中发送消息.
func (mh *MsgHandler) SendMsgToTaskQueue(request iface.RequestI) {
	// 根据ConnID来分配当前的连接应该由哪个worker负责处理
	// 轮询的平均分配法则

	// 得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	log.Debug("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), "to workerID=", workerID)
	// 将请求消息发送给任务队列
	mh.TaskQueue[workerID] <- request
}
