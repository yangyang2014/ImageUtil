package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	// font, err := truetype.Parse(goregular.TTF)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// face := truetype.NewFace(font, &truetype.Options{Size: 48})

	// dc := gg.NewContext(1024, 1024)
	// dc.SetFontFace(face)
	// dc.SetRGB(1, 1, 1)
	// dc.Clear()
	// dc.SetRGB(0, 0, 0)
	// dc.DrawStringAnchored("Hello, world!", 512, 512, 0.5, 0.5)
	// dc.SavePNG("image/gofont.png")
	addTextToImage("C:/Users/yangyangwang/Desktop/testfile/test20201115batchimage/")
}

func addTextToImage(path string) {

	f, err := ioutil.ReadFile(path + "/readme.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}

	var tempImageInfo = string(f)
	imageInfoArray := strings.Split(tempImageInfo, "\n")
	for _, imageInfo := range imageInfoArray {
		fmt.Println("这是新图片")
		atrArr := strings.Fields(imageInfo)
		imageName := atrArr[0]
		dateStr := atrArr[1]
		infoStr := atrArr[2]

		im2, err := gg.LoadImage(path + "image/" + imageName)
		if err != nil {
			panic(err)
		}
		getNewImage(im2, imageName, dateStr, infoStr, path)
	}
}

func getNewImage(im2 image.Image, imageName string, dateStr string, infoStr string, path string) {

	s1 := im2.Bounds().Size()
	imageWidth := s1.X
	imageHeight := s1.Y
	dc := gg.NewContext(imageWidth, imageHeight)
	dc.LoadFontFace("C:/Windows/Fonts/STXINGKA.TTF", 15.0)
	// dc.SetFontFace(face)
	dc.SetRGB(1, 0, 0)

	dc.Clear() //使用颜色填充
	dc.DrawImage(im2, 0, 0)
	dc.SetRGB(0, 0, 0)
	posx := float64(imageWidth) - 10.0
	posy := 10.0
	words := dateStr + " " + infoStr //待绘制的文字空格就表示换行
	lineSpace := 1.5                 //行间距 0.5就表示0.5倍
	fontwidth := 0.0                 //行文字宽度
	//gg.AlignRight 右对齐 是在指定位置(fontwidth+posx,posy)左侧的文本靠右对齐，而不是靠着图片的右侧区域
	//gg.AlignLeft  左对齐 是在指定位置(fontwidth+posx,posy)右侧的文本靠左对齐
	dc.DrawStringWrapped(words, posx, posy, 0, 0, fontwidth, lineSpace, gg.AlignRight)
	dc.SavePNG(path + "/out/" + imageName + ".png")

}
