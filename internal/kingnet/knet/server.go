// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import (
	"fmt"
	"net"
	"time"

	"KingNet/internal/kingnet/iface"
	"KingNet/internal/pkg/log"
)

// Server implements the ServerI interface
// It is the main entry point for the KingNet server.
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
	msgHandler iface.MsgHandlerI
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

		// 1 get the addr from the server
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Errorf("[ERROR] failed to resolve TCP address: %s", err.Error())
			return
		}

		// 2 listener the tcp port
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Errorf("[ERROR] failed to listen on TCP address: %s", err.Error())
			return
		}
		log.Infof("[SUCCESS] server start successfully at %s: %d.", s.IP, s.Port)
		var cid uint32 = 0
		// 3 start the tcp listener
		for {
			connection, err := listener.AcceptTCP()
			if err != nil {
				log.Errorf("[ERROR] listener failed to accept TCP connection: %s, remote addr: %s", err.Error(), connection.RemoteAddr().String())
				continue
			}
			// the number of connections is over the maximum number of connections, decline the connection
			if s.ConnMgr.Len() >= ServerOption.MaxConnections {
				connection.Close()
				continue
			}

			log.Infof("[INFO] server accepted connection from :%s with connection id %d", connection.RemoteAddr().String(), cid)
			dealConn := NewConnection(s, connection, cid, s.msgHandler)
			cid++

			// 4 start the connection handler
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	log.Infof("[STOP] Zinx server , name ", s.Name)

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
		Name:       ServerOption.Name,
		IPVersion:  "tcp4",
		IP:         ServerOption.Host,
		Port:       ServerOption.TcpPort,
		msgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
}

// GetConnMgr returns the connection manager of the server.
func (s *Server) GetConnMgr() iface.ConnManagerI {
	return s.ConnMgr
}

// SetOnConnStart sets the OnConnStart hook function.
func (s *Server) SetOnConnStart(hookFunc func(iface.ConnectionI)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop sets the OnConnStop hook function.
func (s *Server) SetOnConnStop(hookFunc func(iface.ConnectionI)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart Hook.
func (s *Server) CallOnConnStart(conn iface.ConnectionI) {
	if s.OnConnStart != nil {
		log.Info("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop Hook.
func (s *Server) CallOnConnStop(conn iface.ConnectionI) {
	if s.OnConnStop != nil {
		log.Info("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

// AddRouter adds a router to the server.
func (s *Server) AddRouter(msgId uint32, router iface.RouterI) {
	s.msgHandler.AddRouter(msgId, router)
}
