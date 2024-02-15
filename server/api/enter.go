package api

import "gin-vue-blog/server/api/settings_api"

type Group struct {
	SettingApi settings_api.SettingsApi
}

// GroupApp 实例化对象
var GroupApp = new(Group)
