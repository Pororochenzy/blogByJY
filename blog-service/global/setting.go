package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppsettingS
	DatabaseSetting *setting.DatabaseSettingS
)
var (
	DBEngine *gorm.DB
)
