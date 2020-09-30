package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		for i := 0; i < 10; i++ {
			//robotgo.MouseClick("left", true)
			robotgo.TypeStr("い")
			time.Sleep(time.Second * 10)
			robotgo.KeyTap("enter")
		}
		robotgo.KeyTap("enter")
	}
}
