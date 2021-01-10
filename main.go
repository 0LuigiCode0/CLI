package main

import (
	"fmt"

	"github.com/0LuigiCode0/CLI/core"
)

func main() {
	dev := core.Dev().OnActive(func() {
		fmt.Println("dev active")
	})

	l := core.Layout().
		OnCreate(func() {
			fmt.Println("layout create")
		}).
		SetComponents(dev)
	l.AddEvent(core.KeyA, func() {
		fmt.Println("its a")
	})

	app, err := core.InitApp(l)
	if err != nil {
		fmt.Println(err)
		return
	}
	app.Start()
	return
}
