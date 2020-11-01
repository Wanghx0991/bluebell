package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//使用viper管理config.yaml

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.AddConfigPath(".") //指定查找配置文件的路径(使用相对路径)
	err = viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig Failed err = %s",err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改..")
	})
	return nil
}