package beegotest

import (
//	"strings"
)

var (
	AppDeveloper string
)

func SetAppName(appname string) {
	AppName = appname
}

func GetAppName() string {
	if "" != AppName {
		return AppName
	}

	return "DefaultApp"
}
