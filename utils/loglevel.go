package utils

import (
	"github.com/gogf/gf/g/os/glog"
	"strings"
)

func GetLogLevel(s string) int {
	s = strings.ToUpper(s)
	switch s {
	case "DEV":
		return glog.LEVEL_DEV
	case "PROD":
		return glog.LEVEL_PROD
	case "TEST":
		return glog.LEVEL_DEBU | glog.LEVEL_PROD
	}
	return glog.LEVEL_PROD
}
