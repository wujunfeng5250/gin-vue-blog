package settings_api

import (
	"fmt"
	"gin-vue-blog/server/config"
	"gin-vue-blog/server/global"
	"gin-vue-blog/server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.Mysql
	err := c.ShouldBindJSON(cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	fmt.Println("before", global.Config)
	global.Config.Mysql = cr
	fmt.Println("after", global.Config)
}
