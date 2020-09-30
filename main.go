package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		for i := 0; i < 10; i++ {
			//robotgo.MouseClick("left", true)
			robotgo.TypeStr("あい")
			time.Sleep(time.Second * 2)
			robotgo.KeyTap("enter")
		}
		robotgo.KeyTap("enter")
	}
}
