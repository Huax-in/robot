package hook

import (
	"fmt"
	hook "github.com/robotn/gohook"
	// "time"
	"robot/data"
	"robot/jutil"
	"robot/mouse"
	"robot/wins"
)

func Run() {
	go hooks()
}

func hooks() {
	EvChan := hook.Start()
	defer hook.End()
	fmt.Println("hook start")

	var stop_f8 chan int = make(chan int)
	var stop_f10 chan int = make(chan int)
	// 监听
	for ev := range EvChan {
		// if ev.Rawcode == 27 { //esc
		// 	// End()
		// }
		event_f8(stop_f8, ev)
		event_f9(ev)
		event_f10(stop_f10, ev)
	}
}

func event_f8(stop chan int, ev hook.Event) {
	if ev.Rawcode == 119 && ev.Kind == hook.KeyUp { //f8
		var state = wins.Main.GetStateText()
		if state == data.STATE_RUN+data.Main.Mouse {
			wins.Main.SetStateText(data.STATE_STOP)
			fmt.Printf("hook: %v\n", ev)
			stop <- 1
		} else if state == data.STATE_STOP {
			fmt.Printf("hook: %v\n", ev)
			wins.Main.SetStateText(data.STATE_RUN + data.Main.Mouse)
			go mouse.Clicks(stop, data.Main.Mouse, wins.Main.GetTimes())
		}
	}
}

var evs []hook.Event = make([]hook.Event, 0, 100)

func event_f9(ev hook.Event) {
	if ev.Rawcode == 120 && ev.Kind == hook.KeyUp { //f9
		if wins.Main.GetStateText() == data.STATE_STOP {
			wins.Main.SetStateText(data.STATE_RECORD)
			evs = evs[0:0]
		} else if wins.Main.GetStateText() == data.STATE_RECORD {
			wins.Main.SetStateText(data.STATE_STOP)
			go jutil.ToFile(evs, func(file string) {
				wins.Main.ReflushFileList()
				wins.Main.SetFile(file)
			})
		}
	}
	if wins.Main.GetStateText() == data.STATE_RECORD {
		evs = append(evs, ev)
	}
}

func event_f10(stop chan int, ev hook.Event) {
	if ev.Rawcode == 121 && ev.Kind == hook.KeyUp { //f10
		if wins.Main.GetStateText() == data.STATE_STOP {
			var evlist = jutil.ToStruct(wins.Main.GetFile())
			if len(evlist) != 0 {
				wins.Main.SetStateText(data.STATE_REPLAY)
				go Replay(stop, evlist)
			}
		} else if wins.Main.GetStateText() == data.STATE_REPLAY {
			wins.Main.SetStateText(data.STATE_STOP)
			stop <- 1
		}
	}
}
