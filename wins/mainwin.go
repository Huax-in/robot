package wins

import (
	"github.com/lxn/win"
	"log"
	"robot/mwd"

	// "fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"robot/data"
	"robot/jutil"
	"strconv"
	"strings"
)

type MainWin struct {
	window *mwd.MyWindow
	state  *walk.Label
	times  *walk.ComboBox
	files  *walk.ComboBox
}

var Main *MainWin

func init() {
	Main = &MainWin{
		window: new(mwd.MyWindow),
		state:  new(walk.Label),
		times:  new(walk.ComboBox),
		files:  new(walk.ComboBox),
	}

	err := MainWindow{
		AssignTo: &Main.window.MainWindow,
		Title:    "Repeater",
		Size:     Size{Width: 300, Height: 240},
		Layout: Grid{
			Columns: 6,
		}, // VBox  HBox
		DataBinder: DataBinder{
			DataSource: data.Main,
			AutoSubmit: true,
			OnSubmitted: func() {
				// fmt.Println(data.Main)
			},
		},
		OnSizeChanged: func() {
			if win.IsIconic(Main.window.Handle()) {
				Main.window.Hide()
				Main.window.Ni.SetVisible(true)
			}
		},
		Children: []Widget{
			Label{
				Text:       "启动/停止连点： F8",
				ColumnSpan: 6,
			},
			Label{
				Text:       "启动/停止录制： F9",
				ColumnSpan: 6,
			},
			Label{
				Text:       "启动/停止回放： F10",
				ColumnSpan: 6,
			},
			Label{
				Text:       "连点键位：",
				ColumnSpan: 2,
			},
			RadioButtonGroup{
				DataMember: "Mouse",
				Buttons: []RadioButton{
					RadioButton{
						Name:       "mouse",
						Text:       "鼠标左键",
						Value:      data.MOUSE_LEFT,
						ColumnSpan: 2,
					},
					RadioButton{
						Name:       "mouse",
						Text:       "鼠标右键",
						Value:      data.MOUSE_RIGHT,
						ColumnSpan: 2,
					},
				},
			},
			Label{
				Text:       "点击频率：",
				ColumnSpan: 2,
			},
			ComboBox{
				AssignTo:   &Main.times,
				Value:      "1",
				Model:      []string{"1", "10", "50", "100"},
				ColumnSpan: 2,
			},
			Label{
				Text:       "次/秒",
				ColumnSpan: 1,
			},
			Label{
				Text:       "脚本：",
				ColumnSpan: 2,
			},
			ComboBox{
				AssignTo: &Main.files,
				// Value:      "",
				// Model:      []string{},
				ColumnSpan: 3,
			},
			Label{
				Text:       "运行状态：",
				ColumnSpan: 2,
			},
			Label{
				AssignTo:   &Main.state,
				Text:       data.STATE_STOP,
				ColumnSpan: 4,
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}
	Main.window.HWnd = Main.window.Handle()
	Main.window.AddNotifyIcon()
	Main.window.SetSizePersistent(false)
	Main.window.SetMaximizeBox(false)

	Main.ReflushFileList()
	Main.files.SetCurrentIndex(0)
}

func (main *MainWin) Run() {
	main.window.Run()
}

func (main *MainWin) SetStateText(text string) {
	main.state.SetText(text)
}

func (main *MainWin) GetStateText() string {
	return main.state.Text()
}

func (main *MainWin) GetTimes() int {
	times, _ := strconv.Atoi(Main.times.Text())
	return times
}

func (main *MainWin) ReflushFileList() {
	Main.files.SetModel(jutil.GetList())
}

func (main *MainWin) GetFile() string {
	return Main.files.Text()
}

func (main *MainWin) SetFile(file string) {
	value := Main.files.Model()
	for i, str := range value.([]string) {
		if strings.EqualFold(str, file) {
			Main.files.SetCurrentIndex(i)
			break
		}
	}
}
