package settings_api

import (
	"gin-vue-blog/server/global"
	"gin-vue-blog/server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Mysql, c)
}
