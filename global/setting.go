package global

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/setting"
	"github.com/jinzhu/gorm"
)

// 声明全局变量
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
)
