package knet

var (
	ServerOption *Options
)

type Options struct {
	Host    string
	TcpPort int
	Name    string
	Version string

	MaxPacketSize    uint32
	MaxConnections   int
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	MaxMsgChanLen    uint32
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
