package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		//robotgo.MouseClick("left", true)
		robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
		robotgo.ScrollMouse(10, "up")
		time.Sleep(time.Second * 10)
	}
}
