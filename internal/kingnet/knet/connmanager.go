package knet

import (
	"fmt"
	"github.com/pkg/errors"
	"go.mod/internal/kingnet/iface"
	"go.mod/internal/pkg/log"
	"sync"
)

// ConnManager 连接管理器
type ConnManager struct {
	connections map[uint32]iface.ConnectionI //管理的连接信息
	connLock    sync.RWMutex                 //读写连接的读写锁
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.ConnectionI),
	}
}

var _ iface.ConnManagerI = &ConnManager{}

// Add 添加链接
func (connMgr *ConnManager) Add(conn iface.ConnectionI) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn连接添加到ConnMananger中
	connMgr.connections[conn.GetConnID()] = conn

	log.Infof("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

// Remove 删除连接
func (connMgr *ConnManager) Remove(conn iface.ConnectionI) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connections, conn.GetConnID())

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", connMgr.Len())
}

// Get 利用ConnID获取链接
func (connMgr *ConnManager) Get(connID uint32) (iface.ConnectionI, error) {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Len 获取当前连接
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// ClearConn 清除并停止所有连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range connMgr.connections {
		//停止连接
		conn.Stop()
		//删除连接
		delete(connMgr.connections, connID)
	}

	log.Infof("Clear All Connections successfully: conn num = %d", connMgr.Len())
}
