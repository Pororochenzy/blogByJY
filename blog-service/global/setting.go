package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppsettingS
	DatabaseSetting *setting.DatabaseSettingS
	JwtSetting      *setting.JwtSettingS
	EmailSetting    *setting.EmailSettingS
)
var (
	DBEngine *gorm.DB
)

var (
	Logger *logger.Logger
)
