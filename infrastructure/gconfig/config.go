package gconfig

import "github.com/spf13/viper"

var Parameters configParameters

// Load 获取配置参数
func Load() error {
	//表示 先预加载匹配的环境变量
	viper.AutomaticEnv()
	// 从yaml文件获取nacos配置
	vconfig := viper.New()
	// 添加读取的配置文件路径
	vconfig.AddConfigPath("./infrastructure/gconfig/")
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	// 读取yaml
	err := vconfig.ReadInConfig()
	if err != nil {
		return err
	}
	// 转译yaml文件
	if err = vconfig.Unmarshal(&Parameters); err != nil {
		return err
	}
	return nil
}

// configParameters 项目配置参数
type configParameters struct {
	Mode       string           `mapstructure:"mode"`
	DB         dbConfig         `mapstructure:"dbConfig"`
	HttpServer httpServerConfig `mapstructure:"httpServerConfig"`
	ZapLog     zapLogConfig     `zapLogConfig:"mysqlConfig"`
}

// dbConfig 数据库配置
type dbConfig struct {
	SqliteUrl string `mqpstructure:"sqliteUrl"`
}

// httpConfig http server 配置
type httpServerConfig struct {
	ListeningPort string `mapstructure:"listeningPort"`
	CertFilePath  string `mapstructure:"certFilePath"`
	KeyFilePath   string `mapstructure:"keyFilePath"`
}

// zapLogConfig 日志相关配置
type zapLogConfig struct {
	InfoLogConfig infoLogConfig `mapstructure:"infoLogConfig"`
	ErrLogConfig  errLogConfig  `mapstructure:"infoLogConfig"`
}

type logConfig struct {
	LogPath string `mapstructure:"infoLogPath"`

	MaxSize    int  `mapstructure:"maxSize"`
	MaxBackups int  `mapstructure:"maxBackups"`
	MaxAge     int  `mapstructure:"maxAge"`
	Compress   bool `mapstructure:"compress"`
}

type infoLogConfig struct {
	logConfig
}

type errLogConfig struct {
	logConfig
}
