package main

import (
	"github.com/jboelter/notificator"
)

var notify *notificator.Notificator

func main() {

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "My test App",
	})

	notify.PushWithIcon("title", "text", "/home/user/icon.png")
}
