package helpers

import (
	"image"
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"
)

var width, height = 256, 256

func ImageViewer(image image.Image) {
	gl.StartDriver(func(driver gxui.Driver) {
		theme := dark.CreateTheme(driver)
		img := theme.CreateImage()
		img.SetTexture(driver.CreateTexture(image, 1.0))
		window := theme.CreateWindow(width, height, "Images")
		window.AddChild(img)
		window.OnClose(driver.Terminate)
	})
}

