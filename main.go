package main

import (
	"github.com/0LuigiCode0/CLI/core"
)

func main() {
	app := core.InitApp()
	l := core.NewLayout()
	app.Window().SetLayout(l)
	app.Start()
	return
}
