package iface

import "net"

type ConnectionI interface {
	// Start starts the connection
	Start()
	// Stop stops the connection
	Stop()
	// GetTCPConnection returns the underlying TCP connection
	GetTCPConnection() *net.TCPConn
	// GetConnID returns the connection ID
	GetConnID() uint32
	// RemoteAddr returns the remote address
	RemoteAddr() net.Addr

	// SendMsg sends a message to the remote peer
	SendMsg(msgId uint32, data []byte) error
	// SendBuffMsg sends a buffer message to the remote peer
	SendBuffMsg(msgId uint32, data []byte) error

	// SetProperty sets a property
	SetProperty(key string, value interface{})
	// GetProperty returns a property
	GetProperty(key string) (interface{}, error)
	// RemoveProperty removes a property
	RemoveProperty(key string)
}
