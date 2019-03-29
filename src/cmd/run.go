package cmd

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/fujiwara/go-zabbix-get/zabbix"
)

//RunAgent run a agent
func RunAgent(addr, port string) error {
	err := zabbix.RunAgent(addr+":"+port, func(key string) (string, error) {
		switch key {
		case "agent.go.ping":
			logs.Debug("key:", key, "=======>", "1")
			return "1", nil
		case "cpu.model":
			logs.Debug("key:", key, "=======>", GetCPUInfo())
			return GetCPUInfo(), nil
		case "system.info":
			logs.Debug("key:", key, "=======>", GetSystemDescription())
			return GetSystemDescription(), nil
		case "mem.total":
			logs.Debug("key:", key, "=======>", GetMem("total"))
			return GetMem("total"), nil
		case "mem.used":
			logs.Debug("key:", key, "=======>", GetMem("used"))
			return GetMem("used"), nil
		case "mem.usedper":
			logs.Debug("key:", key, "=======>", GetMem("usedper"))
			return GetMem("usedper"), nil
		default:
			return "", errors.New("not supported")
		}
	})
	if err != nil {
		return err
	}
	return nil
}
