package flag

import (
	"gin-vue-blog/server/global"
	"gin-vue-blog/server/models"
)

func MakeMigrations() {
	var err error
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.ArticleModel{},
			&models.UserCollectModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表成功")
}
