package iface

type ServerI interface {

	// Start starts the server
	Start()
	// Stop stops the server
	Stop()
	// Serve define the methods to do something
	Serve()

	// GetConnMgr returns the Connection Manager
	GetConnMgr() ConnManagerI

	// SetOnConnStart sets the function to be called when a new connection is established
	SetOnConnStart(func(ConnectionI))
	// SetOnConnStop define the methods to do something when a connection is stopped
	SetOnConnStop(func(ConnectionI))
	// CallOnConnStart calls the function set with SetOnConnStart
	CallOnConnStart(conn ConnectionI)
	// CallOnConnStop calls the function set with SetOnConnStop
	CallOnConnStop(conn ConnectionI)
}
