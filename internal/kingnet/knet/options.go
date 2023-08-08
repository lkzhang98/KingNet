package knet

var (
	ServerOption *Options
)

// Options 存放knet的配置
type Options struct {
	Host    string // 监听的地址
	TcpPort int    // 监听的端口
	Name    string // 应用名称
	Version string // 版本信息

	MaxPacketSize    uint32 // 接受包的最大尺寸
	MaxConnections   int    // 最大连接数
	WorkerPoolSize   uint32 // 工作池的最大数
	MaxWorkerTaskLen uint32 // 每个worker的最大任务数
	MaxMsgChanLen    uint32 // 最大的消息channel的长度
}

func NewServerOptions() *Options {
	return &Options{
		Host:    "0.0.0.0",
		TcpPort: 8888,
		Name:    "KingNet",
		Version: "1.0",
	}
}

func Init(opts *Options) {
	ServerOption = &Options{
		Host:             opts.Host,
		TcpPort:          opts.TcpPort,
		Name:             opts.Name,
		Version:          opts.Version,
		MaxPacketSize:    opts.MaxPacketSize,
		MaxConnections:   opts.MaxConnections,
		WorkerPoolSize:   opts.WorkerPoolSize,
		MaxWorkerTaskLen: opts.MaxWorkerTaskLen,
		MaxMsgChanLen:    opts.MaxMsgChanLen,
	}
}
