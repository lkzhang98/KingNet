// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import (
	"reflect"
	"testing"

	"KingNet/internal/kingnet/iface"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddRouter(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	type args struct {
		msgId  uint32
		router iface.RouterI
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.AddRouter(tt.args.msgId, tt.args.router)
		})
	}
}

func TestServer_CallOnConnStart(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	type args struct {
		conn iface.ConnectionI
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.CallOnConnStart(tt.args.conn)
		})
	}
}

func TestServer_CallOnConnStop(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	type args struct {
		conn iface.ConnectionI
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.CallOnConnStop(tt.args.conn)
		})
	}
}

func TestServer_GetConnMgr(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
		want   iface.ConnManagerI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			if got := s.GetConnMgr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConnMgr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Serve(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.Serve()
		})
	}
}

func TestServer_SetOnConnStart(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	type args struct {
		hookFunc func(iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.SetOnConnStart(tt.args.hookFunc)
		})
	}
}

func TestServer_SetOnConnStop(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	type args struct {
		hookFunc func(iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.SetOnConnStop(tt.args.hookFunc)
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.Start()
		})
	}
}

func TestServer_Stop(t *testing.T) {
	type fields struct {
		Name        string
		IPVersion   string
		IP          string
		Port        int
		msgHandler  iface.MsgHandlerI
		ConnMgr     iface.ConnManagerI
		OnConnStart func(conn iface.ConnectionI)
		OnConnStop  func(conn iface.ConnectionI)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Name:        tt.fields.Name,
				IPVersion:   tt.fields.IPVersion,
				IP:          tt.fields.IP,
				Port:        tt.fields.Port,
				msgHandler:  tt.fields.msgHandler,
				ConnMgr:     tt.fields.ConnMgr,
				OnConnStart: tt.fields.OnConnStart,
				OnConnStop:  tt.fields.OnConnStop,
			}
			s.Stop()
		})
	}
}
