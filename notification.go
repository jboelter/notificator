package notificator

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Options for creating a new notificator
type Options struct {
	DefaultIcon string
	AppName     string
}

type notifier interface {
	push(title string, text string, iconPath string)
}

// The Notificator type
type Notificator struct {
	notifier    notifier
	defaultIcon string
}

// PushWithIcon displays a notification with an icon, title and message
func (n Notificator) PushWithIcon(title string, text string, iconPath string) {

	icon := n.defaultIcon

	if iconPath != "" {
		icon = iconPath
	}

	n.notifier.push(title, text, icon)
}

// Push displays a notification with a title, message and default icon
func (n Notificator) Push(title string, text string) {

	n.notifier.push(title, text, n.defaultIcon)
}

type osxNotificator struct {
	AppName string
}

func (o osxNotificator) push(title string, text string, iconPath string) {
	args := make([]string, 0, 7)
	if iconPath != "" {
		args = append(args, "--image", iconPath)
	}
	args = append(args, "-n", o.AppName, "-m", title, text)

	exec.Command("growlnotify", args...).Run()
}

type linuxNotificator struct{}

func (l linuxNotificator) push(title string, text string, iconPath string) {
	args := make([]string, 0, 4)
	if iconPath != "" {
		args = append(args, "-i", iconPath)
	}
	args = append(args, title, text)

	fmt.Println(args)

	exec.Command("notify-send", args...).Run()
}

type windowsNotificator struct{}

func (w windowsNotificator) push(title string, text string, iconPath string) {
	args := make([]string, 0, 5)
	if iconPath != "" {
		args = append(args, "/i:", iconPath)
	}
	args = append(args, "/t:", title, text)

	exec.Command("growlnotify", args...).Run()
}

// New creates a new Notificator with the specified Options
func New(o Options) *Notificator {

	var notifier notifier

	switch runtime.GOOS {

	case "darwin":
		notifier = osxNotificator{AppName: o.AppName}
	case "linux":
		notifier = linuxNotificator{}
	case "windows":
		notifier = windowsNotificator{}

	}

	return &Notificator{notifier: notifier, defaultIcon: o.DefaultIcon}
}
