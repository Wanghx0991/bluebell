package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string  `mapstructure:"name"`
	Mode string  `mapstructure:"mode"`
	Port string  `mapstructure:"port"`

	StartTime string `mapstructure:"start_time"`
	MachineID int64 `mapstructure:"machine_id"`

	*LogConfig   `mapstructure:"log"`

	*MySQLConfig `mapstructure:"mysql"`

	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Db  int `mapstructure:"db"`
	Password string `mapstructure:"password"`
}
// 全局变量,用来保存程序的所有配置信息
var Conf = new(AppConfig)
//使用viper管理config.yaml

//定义了结构体后,程序启动还是使用viper加载yaml信息,加载完后反序列化到结构体变量里,
//后续在程序中使用配置信息时,直接使用结构体即可
func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件
	err = viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig Failed err = %s",err)
		return
	}
	//把读取到的配置信息反序列化到结构体conf中
	if err := viper.Unmarshal(Conf);err != nil{
		fmt.Printf("viper.Unmarshal failed, err %v \n",err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改..")
		if err := viper.Unmarshal(Conf);err != nil{
			fmt.Printf("viper.Unmarshal failed, err %v \n",err)
		}
	})
	return nil
}