package setting

import (
	"github.com/go-ini/ini"
	"github.com/prometheus/common/log"
)

var cfg *ini.File

// App struct
type App struct {
	Debug      string
	ListenIP   string
	ListenPort string
	LogFile    string
}

// AppSetting app
var AppSetting = &App{}

//LoadConfig load config file
func LoadConfig(config string) error {
	var err error
	cfg, err = ini.Load(config)
	if err != nil {
		return err
	}
	mapTo("app", AppSetting)
	return nil
}

//mapTo map to config
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorln(err)
	}
}
