package main

import "github.com/0xAX/notificator"


func SendNotification(result []string) {
	var notify *notificator.Notificator
	notify = notificator.New(notificator.Options{
		DefaultIcon: Icon,
		AppName:	"Zabbix",
	})

	for _, val := range result {
		notify.Push("Server Alert", val, Icon, notificator.UR_CRITICAL)
	}

}
