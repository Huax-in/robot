package mouse

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func Clicks(stop chan int, key string, times float64) {
	fmt.Println("mouse click start " + key)
loop:
	for {
		select {
		case <-stop:
			fmt.Println("mouse click stop")
			break loop
		default:
		}
		time.Sleep(time.Duration(1000/times) * time.Millisecond)
		robotgo.MouseClick(key, false)
	}
}

func Move(x, y int) {
	robotgo.MoveMouseSmooth(x, y, 0.0, 0.0, 10)
}

func Click(key uint16) {
	if key == 1 {
		robotgo.MouseClick("left", false)
	} else if key == 2 {
		robotgo.MouseClick("right", false)
	}
}

func Down(key uint16) {
	if key == 1 {
		robotgo.MouseToggle("down", "left")
	} else if key == 2 {
		robotgo.MouseToggle("down", "right")
	}
}

func Up(key uint16) {
	if key == 1 {
		robotgo.MouseToggle("up", "left")
	} else if key == 2 {
		robotgo.MouseToggle("up", "right")
	}
}
