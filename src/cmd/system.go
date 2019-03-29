package cmd

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/host"
)

//GetSystemDescription system info
func GetSystemDescription() string {
	v, _, v2, err := host.PlatformInformation()
	if err != nil {
		logs.Error(err)
		return ""
	}
	desc := fmt.Sprintf("%v %v", v, v2)
	return desc
}
