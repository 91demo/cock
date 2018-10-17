package main

import (
	"flag"

	"github.com/eagle-1949/cock/dinghook"
)

var accesstoken = flag.String("accesstoken", "access_token", "钉钉机器人webhook的access_token")
var message = flag.String("message", "zabbix告警发送的消息", "接受zabbix发送的告警消息，在zabbix action中设置")
var ding *dinghook.DingTalk

func main() {
	flag.Parse()
	ding = dinghook.NewDingTalk()
	ding.Send(*accesstoken, *message)

}
