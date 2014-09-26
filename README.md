notificator
===========================

Desktop notification with golang for:

  * Windows with `growlnotify`;
  * Mac OS X with `growlnotify`;
  * Linux with `notify-send` for gnome and `kdialog` for kde.

usage
------

```go
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
```

todo
-----

  * Add more options for different notificators.

contribution
------------

  * Fork;
  * Make changes;
  * Send pull request;
  * Thank you.

authors
----------

[@0xAX](https://twitter.com/0xAX)
joshuaboelter
