package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		robotgo.MouseClick("left", true)
		time.Sleep(time.Second * 10)
	}
}
