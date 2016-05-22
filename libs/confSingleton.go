package libs

import (
	"log"
)

var oneSingleton *ConfigFile

//单例模式，保证整个服务器只存在一份map数据
func NewSingleton() (*ConfigFile, error) {
	if oneSingleton == nil {
		// 创建并获取一个 ConfigFile 对象，以进行后续操作
		// 文件名支持相对和绝对路径
		cfg, err := LoadConfigFile("database/user.ini", "database/sets.ini")
		if err != nil {
			log.Fatalf("无法加载配置文件：%s", err)
			return nil, err
		}
		oneSingleton = cfg
	}
	return oneSingleton, nil
}
