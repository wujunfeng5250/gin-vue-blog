package routers

import (
	"gin-vue-blog/server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.GroupApp.SettingApi
	router.GET("settings", settingsApi.SettingsInfoView)
	router.PUT("settings_update", settingsApi.SettingsInfoUpdateView)
}
