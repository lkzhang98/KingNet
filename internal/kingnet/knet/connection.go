package knet

import (
	"errors"
	"go.mod/internal/kingnet/iface"
	"go.mod/internal/pkg/log"
	"io"
	"net"
	"sync"
)

// Connection 用来处理连接.
type Connection struct {
	TcpServer iface.ServerI
	// Conn 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// ConnID 当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	// isClosed 当前连接的关闭状态
	isClosed bool

	// 该连接的处理方法api
	MsgHandler iface.MsgHandlerI

	// 告知该链接已经退出/停止的channel
	ExitBuffChan chan bool

	// 无缓冲管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte
	// 有缓冲管道，用于读、写两个goroutine之间的消息通信
	msgBuffChan chan []byte

	property map[string]interface{}
	// 保护链接属性修改的锁
	propertyLock sync.RWMutex
}

func NewConnection(server iface.ServerI, conn *net.TCPConn, connID uint32, msgHandler iface.MsgHandlerI) *Connection {
	c := &Connection{
		TcpServer:    server,
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		MsgHandler:   msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, ServerOption.MaxMsgChanLen),
		property:     make(map[string]interface{}),
	}
	// 将新创建的Conn添加到链接管理中
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

var _ iface.ConnectionI = &Connection{}

// StartReader 处理conn读数据的Goroutine.
func (c *Connection) StartReader() {
	log.Debug("Reader Goroutine is running")
	defer log.Debug(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 读取我们最大的数据到buf中
		dp := NewDataPack()

		// 读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			log.Errorf("[ERROR] connection %d read msg head error:%v", c.ConnID, err)
			c.ExitBuffChan <- true
			continue
		}

		// 拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			log.Errorf("[ERROR] unpack data error: %v from %s.", err, c.RemoteAddr().String())
			c.ExitBuffChan <- true
			continue
		}

		// 根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				log.Errorf("[ERROR] connection %d read msg data error: %v", c.Conn, err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		// 得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		// 从路由Routers 中找到注册绑定Conn的对应Handle
		if ServerOption.WorkerPoolSize > 0 {
			// 已经启动工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			// 从绑定好的消息和对应的处理方法中执行对应的Handle方法
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// Start 启动连接，让当前连接开始工作.
func (c *Connection) Start() {
	//开启处理该链接读取到客户端数据之后的请求业务
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()
	// 2 开启用于写回客户端数据流程的Goroutine
	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)
	for {
		select {
		case <-c.ExitBuffChan:
			// 得到退出消息，不再阻塞
			return
		}
	}
}

// Stop 停止连接，结束当前连接状态M.
func (c *Connection) Stop() {
	log.Infof("Conn Stop()...ConnID = ", c.ConnID)
	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//==================
	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	c.TcpServer.CallOnConnStop(c)
	//==================

	// 关闭socket链接
	c.Conn.Close()
	// 关闭Writer
	c.ExitBuffChan <- true

	// 将链接从连接管理器中删除
	c.TcpServer.GetConnMgr().Remove(c)

	// 关闭该链接全部管道
	close(c.ExitBuffChan)
	close(c.msgBuffChan)
}

// GetTCPConnection 从当前连接获取原始的socket TCPConn.
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接ID.
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端地址信息.
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送不带缓冲的信息.
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	// 将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		log.Errorf("Pack error msg id = , Connection Id = ", msgId, c.ConnID)
		return errors.New("Pack error msg ")
	}
	// 写回客户端
	// 将之前直接回写给conn.Write的方法 改为 发送给Channel 供Writer读取
	c.msgChan <- msg

	return nil
}

// SendBuffMsg 发送带有缓存的消息.
func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send buff msg")
	}
	// 将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		log.Errorf("Pack error msg id = , Connection ID = ", msgId, c.ConnID)
		return errors.New("Pack error msg ")
	}

	// 写回客户端
	c.msgBuffChan <- msg

	return nil
}

// StartWriter 将消息写回客户端.
func (c *Connection) StartWriter() {
	log.Info("[Writer Goroutine is running]")
	defer log.Info(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			// 有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				log.Errorf("Send Data error:%v , Connection ID = %d,  Conn Writer exit!", err, c.ConnID)
				return
			}
			// 针对有缓冲channel需要些的数据处理
		case data, ok := <-c.msgBuffChan:
			if ok {
				// 有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					log.Errorf("Send Buff Data error:%v , Connection ID = %d,  Conn Writer exit!", err, c.ConnID)
					return
				}
			} else {
				log.Errorf("msgBuffChan is Closed, Connection Id = %d", c.ConnID)

				break
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

// SetProperty 为当前连接添加属性.
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

// GetProperty 获取链接属性.
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}

	return nil, errors.New("no property found")
}

// RemoveProperty 移除链接属性.
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
