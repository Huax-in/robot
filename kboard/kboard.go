package kboard

import (
	"github.com/go-vgo/robotgo"
	"robot/data"
	"time"
)

//robotgo.TypeStr("だんしゃり我们的繁复", 100.0)
func Down(key int) {
	var kstr, ok = data.Keys[key]
	if ok {
		robotgo.KeyToggle(kstr, "down")
	}
	time.Sleep(50 * time.Millisecond)
}

func Up(key int) {
	var kstr, ok = data.Keys[key]
	if ok {
		robotgo.KeyToggle(kstr, "up")
	}
	time.Sleep(50 * time.Millisecond)
}
