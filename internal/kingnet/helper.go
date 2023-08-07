package kingnet

import "KingNet/internal/pkg/log"

var (
	recommendedHomeDir = ".kingnet"

	// defaultConfigName 指定了 miniblog 服务的默认配置文件名.
	defaultConfigName = "kingnet.yaml"
)

func initConfig() {

}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func init
