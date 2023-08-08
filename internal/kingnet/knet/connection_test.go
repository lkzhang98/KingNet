package knet

import (
	"go.mod/internal/kingnet/iface"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestConnection_GetConnID(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			if got := c.GetConnID(); got != tt.want {
				t.Errorf("GetConnID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_GetProperty(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			got, err := c.GetProperty(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProperty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_GetTCPConnection(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *net.TCPConn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			if got := c.GetTCPConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTCPConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_RemoteAddr(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   net.Addr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			if got := c.RemoteAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoteAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_RemoveProperty(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	type args struct {
		key string
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
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.RemoveProperty(tt.args.key)
		})
	}
}

func TestConnection_SendBuffMsg(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	type args struct {
		msgId uint32
		data  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			if err := c.SendBuffMsg(tt.args.msgId, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendBuffMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnection_SendMsg(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	type args struct {
		msgId uint32
		data  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			if err := c.SendMsg(tt.args.msgId, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnection_SetProperty(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	type args struct {
		key   string
		value interface{}
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
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.SetProperty(tt.args.key, tt.args.value)
		})
	}
}

func TestConnection_Start(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.Start()
		})
	}
}

func TestConnection_StartReader(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.StartReader()
		})
	}
}

func TestConnection_StartWriter(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.StartWriter()
		})
	}
}

func TestConnection_Stop(t *testing.T) {
	type fields struct {
		TcpServer    iface.ServerI
		Conn         *net.TCPConn
		ConnID       uint32
		isClosed     bool
		MsgHandler   iface.MsgHandlerI
		ExitBuffChan chan bool
		msgChan      chan []byte
		msgBuffChan  chan []byte
		property     map[string]interface{}
		propertyLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				TcpServer:    tt.fields.TcpServer,
				Conn:         tt.fields.Conn,
				ConnID:       tt.fields.ConnID,
				isClosed:     tt.fields.isClosed,
				MsgHandler:   tt.fields.MsgHandler,
				ExitBuffChan: tt.fields.ExitBuffChan,
				msgChan:      tt.fields.msgChan,
				msgBuffChan:  tt.fields.msgBuffChan,
				property:     tt.fields.property,
				propertyLock: tt.fields.propertyLock,
			}
			c.Stop()
		})
	}
}

func TestNewConnection(t *testing.T) {
	type args struct {
		server     iface.ServerI
		conn       *net.TCPConn
		connID     uint32
		msgHandler iface.MsgHandlerI
	}
	tests := []struct {
		name string
		args args
		want *Connection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnection(tt.args.server, tt.args.conn, tt.args.connID, tt.args.msgHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
