package iface

// DataPackI 定义拆包和封包方法.
type DataPackI interface {
	GetHeadLen() uint32                // 获取包头长度方法
	Pack(msg MessageI) ([]byte, error) // 封包方法
	Unpack([]byte) (MessageI, error)   // 拆包方法
}
