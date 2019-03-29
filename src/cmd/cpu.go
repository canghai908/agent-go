package cmd

import (
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/cpu"
)

//GetCPUInfo get cpu info
func GetCPUInfo() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		logs.Error(err)
		return "Unknown"
	}
	ModelName := "Unknown"
	if len(cpuInfo) > 0 {
		ci := cpuInfo[0]
		ModelName = ci.ModelName
	}
	return ModelName
}
