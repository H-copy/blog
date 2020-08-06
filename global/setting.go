package global

import (
	"blog/pkg/setting"
	"blog/pkg/logger"
)

var (
	ServerSetting *setting.ServerSetting
	AppSetting *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger *logger.Logger
)