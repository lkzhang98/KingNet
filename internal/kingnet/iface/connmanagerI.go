package iface

type ConnManagerI interface {
	// Add adds a connection to the connection manager.
	Add(conn ConnectionI)
	// Remove removes a connection from the connection manager.
	Remove(conn ConnectionI)
	// Get gets a connection from the connection manager.
	Get(connID uint32) (ConnectionI, error)
	// Len returns the number of connections in the connection manager.
	Len() int
	// ClearConn removes all connections from the connection manager.
	ClearConn()
}
