package kingnet

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mod/internal/kingnet/knet"
	"go.mod/internal/pkg/log"
)

var (
	recommendedHomeDir = ".kingnet"

	//
	defaultConfigName = "kingnet.yaml"
)

var prefix string

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("../../configs/")

		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	viper.SetEnvPrefix("KINGNET")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}
	prefix = viper.GetString("runmode") + "."

	// 打印 viper 当前使用的配置文件，方便 Debug.
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool(prefix + "log.disable-caller"),
		DisableStacktrace: viper.GetBool(prefix + "log.disable-stacktrace"),
		Level:             viper.GetString(prefix + "log.level"),
		Format:            viper.GetString(prefix + "log.format"),
		OutputPaths:       viper.GetStringSlice(prefix + "log.output-paths"),
	}
}

func serverOptions() *knet.Options {
	return &knet.Options{
		Host:             viper.GetString(prefix + "server.host"),
		TcpPort:          viper.GetInt(prefix + "server.port"),
		Name:             viper.GetString(prefix + "server.name"),
		Version:          viper.GetString(prefix + "server.version"),
		MaxPacketSize:    viper.GetUint32(prefix + "server.max-packet-size"),
		MaxConnections:   viper.GetInt(prefix + "server.max-connections"),
		WorkerPoolSize:   viper.GetUint32(prefix + "server.worker-pool-size"),
		MaxWorkerTaskLen: viper.GetUint32(prefix + "server.max-worker-task-length"),
		MaxMsgChanLen:    viper.GetUint32(prefix + "server.max-message-channel-length"),
	}
}
