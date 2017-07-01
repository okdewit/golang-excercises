package slices

import (
	"image"
)

func Data(dx, dy int) [][]uint8 {
	var img = make([][]uint8, dx)
	for x := range img {
		img[x] = make([]uint8, dy)
		for y := range img[x] {
			img[x][y] = uint8(x ^ y)
		}
	}
	return img
}

func GetImage() image.Image {
	const (
		dx = 256
		dy = 256
	)
	data := Data(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	return m
}

