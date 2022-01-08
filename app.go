package main

import (
	"robot/hook"
	// "robot/screen"
	"robot/wins"
)

func main() {

	//
	// go screen.Fullshot()
	// 启动按键监听线程
	hook.Run()
	// 启动主窗体
	wins.Main.Run()
}
