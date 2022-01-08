package jutil

import (
	"bufio"
	"encoding/json"

	// "fmt"
	"io"
	"io/ioutil" //io 工具包
	"os"
	"strconv"
	"strings"
	"time"

	hook "github.com/robotn/gohook"
)

const (
	FILE_PATH = "out/"
	FILE_TYPE = ".data"
)

func init() {
	_, err := os.Stat(FILE_PATH) //创建输出目录
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		os.MkdirAll(FILE_PATH, os.ModePerm)
		return
	}
}

func ToFile(evs []hook.Event, callback func(file string)) {
	var fileName = strconv.FormatInt(time.Now().Unix(), 16) + FILE_TYPE
	var path = FILE_PATH + fileName
	var f *os.File

	f, _ = os.Create(path) //创建文件
	defer f.Close()
	for _, evx := range evs {
		js, _ := json.Marshal(evx)
		io.WriteString(f, string(js)+"\n")
	}
	if callback != nil {
		callback(fileName)
	}
}

func ToStruct(file string) []hook.Event {
	var evs []hook.Event
	if strings.EqualFold(file, "") {
		return evs
	}
	f, _ := os.OpenFile(FILE_PATH+file, os.O_RDONLY, 0600)
	defer f.Close()
	buf := bufio.NewScanner(f)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line = strings.TrimSpace(line)
		var ev = &hook.Event{}
		json.Unmarshal([]byte(line), ev)
		evs = append(evs, *ev)
	}

	return evs
}

func GetList() []string {
	var pathlist []string
	files, _ := ioutil.ReadDir(FILE_PATH)
	var path string
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			path = file.Name()
			// HasSuffix: endwith
			if strings.HasSuffix(path, FILE_TYPE) {
				pathlist = append(pathlist, path)
			}
		}
	}
	return pathlist
}
