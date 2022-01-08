package hook

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"robot/kboard"
	"robot/mouse"
)

func Replay(stop chan int, evlist []hook.Event) {
	i := 0
loof:
	for {
		select {
		case <-stop:
			fmt.Println("replay stop")
			break loof
		default:
		}
		aev := evlist[i]
	loos:
		switch aev.Kind {
		case hook.MouseMove:
		case hook.MouseDrag:
			mouse.Move(int(aev.X), int(aev.Y))
			break loos
		case hook.MouseUp:
			mouse.Move(int(aev.X), int(aev.Y))
			mouse.Click(aev.Button)
			break loos
		case hook.MouseHold:
			mouse.Move(int(aev.X), int(aev.Y))
			mouse.Down(aev.Button)
			break loos
		case hook.MouseDown:
			mouse.Move(int(aev.X), int(aev.Y))
			mouse.Up(aev.Button)
			break loos
		case hook.KeyDown:
		case hook.KeyHold:
			kboard.Down(int(aev.Rawcode))
			break loos
		case hook.KeyUp:
			kboard.Up(int(aev.Rawcode))
			break loos
		default:
			// fmt.Println(aev)
		}
		i++
		if i == len(evlist) {
			i = 0
		}
	}
}
