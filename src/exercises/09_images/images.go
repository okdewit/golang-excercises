package main

import (
	"image/color"
	"image"
	"golang.org/x/tour/pic"
)

type Image struct{
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 0, 0}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
