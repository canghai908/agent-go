package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/canghai908/agent-go/cmd"
	"github.com/canghai908/agent-go/setting"
	"github.com/urfave/cli"
)

//VER version
const VER = "0.0.1"

var app *cli.App

func main() {
	app = cli.NewApp()
	app.Version = VER
	app.Name = "agent-go"
	app.Usage = "A zabbix agent written by golang"
	app.Action = runAgent
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
			Value: "agent-go.ini",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func runAgent(c *cli.Context) {
	var config string
	config = "/etc/agent-go/agent-go.ini"
	if c.IsSet("config") {
		config = c.String("config")
	}
	err := setting.LoadConfig(config)
	if err != nil {
		logs.Error(err)
	}
	var debug, listenIP, listenPort, logfilename string
	debug = "1"
	if setting.AppSetting.Debug != "" {
		debug = strings.TrimSpace(setting.AppSetting.Debug)
	}
	logfilename = "/var/log/agent-go.log"
	if setting.AppSetting.LogFile != "" {
		logfilename = strings.TrimSpace(setting.AppSetting.LogFile)
	}
	listenIP = "0.0.0.0"
	if setting.AppSetting.ListenIP != "" {
		listenIP = setting.AppSetting.ListenIP
	}
	listenPort = "10049"
	if setting.AppSetting.ListenPort != "" {
		listenPort = setting.AppSetting.ListenPort
	}
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logfilename+`","level":`+debug+`,"maxlines":0,
		"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	err = cmd.RunAgent(listenIP, listenPort)
	if err != nil {
		logs.Error(err)
	}
}
