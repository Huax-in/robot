package screen

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"

	"github.com/kbinani/screenshot"

	// "image/png"
	"os"
	"time"
)

var imagelist = make([]*image.RGBA, 0, 10)

func Fullshot() {
	var chans = make([]chan int, 0, 10)
	for i := 0; i < 400; i++ {
		time.Sleep(200 * time.Millisecond)
		imagelist = append(imagelist, &image.RGBA{})
		chans = append(chans, make(chan int))
		go shot(i, chans[i])
	}
	for i, achan := range chans {
		fmt.Println(i)
		<-achan
	}
	fmt.Println("shot over")
	createGif()
}

var w, h int = 0, 0

func shot(i int, achan chan int) {
	n := screenshot.NumActiveDisplays()

	if n > 0 {
		bounds := screenshot.GetDisplayBounds(0)
		w = bounds.Dx()
		h = bounds.Dy()
		var err error
		imagelist[i], err = screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
	}
	achan <- 1
}

func createGif() {
	var disposals []byte
	var images []*image.Paletted
	var delays []int
	var chans = make([]chan int, 0, 10)
	fmt.Println("len ", len(imagelist))
	for i := 0; i < len(imagelist); i++ {
		// image.RGBA 转为 image.image
		g := (*imagelist[i]).SubImage(image.Rect(0, 0, w, h))
		// 透明图片需要设置
		disposals = append(disposals, gif.DisposalBackground)
		// 新建调色板
		p := image.NewPaletted(image.Rect(0, 0, w, h), color.Palette(palette.Plan9))

		// 填充image.image到image.Paletted(多线程)
		chans = append(chans, make(chan int))
		go func(achan chan int, p *image.Paletted, r image.Rectangle, src image.Image, sp image.Point, op draw.Op) {
			draw.Draw(p, r, src, sp, op)
			achan <- 1
		}(chans[i], p, p.Bounds(), g, image.Point{0, 0}, draw.Src)

		images = append(images, p)
		delays = append(delays, 20)

		// qrImg.(*image.Paletted).Palette = append(qrImg.(*image.Paletted).Palette,color.RGBA{R:255,G:0,B:0,A:255})
	}

	for i, achan := range chans {
		fmt.Println(i)
		<-achan
	}

	g := &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: -1,
		Disposal:  disposals,
	}
	f, err := os.Create("out/test1.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()
	gif.EncodeAll(f, g)
	imagelist = imagelist[0:0]
	fmt.Println("createGif over")
}
