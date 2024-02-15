package main

import (
	"gin-vue-blog/server/core"
	"gin-vue-blog/server/flag"
	"gin-vue-blog/server/global"
	"gin-vue-blog/server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	// gorm 连接数据库
	global.DB = core.InitGorm()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Info("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
