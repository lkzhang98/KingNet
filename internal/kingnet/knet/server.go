package knet

import (
	"KingNet/internal/kingnet/iface"
	"KingNet/internal/pkg/log"
	"fmt"
	"net"
	"time"
)

// Server implements the ServerI interface
// It is the main entry point for the KingNet server
type Server struct {
	// Name is the name of the server
	Name string
	// IPVersion is the IP version of the server
	IPVersion string
	// IP is the IP address of the server
	IP string
	// Port is the port of the server listening on
	Port int
	// msgHandler defines the server how to handle messages
	msgHandler iface.MsgHandleI
	// ConnMgr defines the connection manager
	ConnMgr iface.ConnManagerI

	// onConnStart is called when a new connection is established
	OnConnStart func(conn iface.ConnectionI)
	// onConnStop is called when a connection is stopped
	OnConnStop func(conn iface.ConnectionI)
}

var _ iface.ServerI = (*Server)(nil)

func (s *Server) Start() {
	log.Infof("[START] server listener at IP: %s, Port:%d.", s.IP, s.Port)
	go func() {
		s.msgHandler.StartWorkerPool()

		//1 get the addr from the server
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Errorf("[ERROR] failed to resolve TCP address: %s", err.Error())
			return
		}

		//2 listener the tcp port
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Errorf("[ERROR] failed to listen on TCP address: %s", err.Error())
			return
		}
		log.Infof("[SUCCESS] server start successfully at %s:%d.", s.IP, s.Port)
		var cid uint32
		cid = 0
		//3 start the tcp listener
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Errorf("[ERROR] listener failed to accept TCP connection: %s, remote addr: %s", err.Error(), conn.RemoteAddr().String())
				continue
			}

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()

}

func (s *Server) Serve() {
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer() *Server {

	return &Server{
		Name:       utils.GlobalObject.Name, //从全局参数获取
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,    //从全局参数获取
		Port:       utils.GlobalObject.TcpPort, //从全局参数获取
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
}

// GetConnMgr returns the connection manager of the server
func (s *Server) GetConnMgr() iface.ConnManagerI {
	return s.ConnMgr
}

// SetOnConnStart sets the OnConnStart hook function
func (s *Server) SetOnConnStart(hookFunc func(iface.ConnectionI)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop sets the OnConnStop hook function
func (s *Server) SetOnConnStop(hookFunc func(iface.ConnectionI)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart Hook
func (s *Server) CallOnConnStart(conn iface.ConnectionI) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop Hook
func (s *Server) CallOnConnStop(conn iface.ConnectionI) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}