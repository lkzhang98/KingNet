package iface

type RequestI interface {
	// GetConnection returns the connection
	GetConnection() ConnectionI
	// GetData returns the data
	GetData() []byte
	// GetMsgID returns the message ID
	GetMsgID() uint32
}
