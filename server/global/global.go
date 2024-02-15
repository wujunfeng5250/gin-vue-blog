package global

import (
	"gin-vue-blog/server/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config //需要一个全局变量保存配置文件
	DB     *gorm.DB
	Log    *logrus.Logger
)
