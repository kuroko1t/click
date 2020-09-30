package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		//robotgo.MouseClick("left", true)
		robotgo.TypeStr("あい")
		time.Sleep(time.Second * 10)
		robotgo.KeyTap("enter")
	}
}
