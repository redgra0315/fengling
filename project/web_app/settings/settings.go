package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 全部变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level     string `mapstructure:"level"`
	FileName  string `mapstructure:"filename"`
	MaxSize   int    `mapstructure:"maxsize"`
	MaxAge    int    `mapstructure:"maxage"`
	MaxBackup int    `mapstructure:"maxbackup"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//viper.SetConfigFile("config")                                                // 指定配置文件
	//viper.SetConfigType("yaml")                                                  // 指定配置文件类型
	viper.SetConfigFile("project/web_app/config.yaml") // 指定查找配置文件的路径
	// 读取配置信息
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	//　反序列到　Conf 变量
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Warn("配置文件进行了改变")
		if err := viper.Unmarshal(Conf); err != nil { // 读取配置信息失败
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
	return
}
