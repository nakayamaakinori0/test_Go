package main

import "golang.org/x/tour/pic"
//import "fmt"
import "image"
import "image/color"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(x * y)
		}
		image[y] = row
	}
	return image
}

type imageSize struct {
	dx int
	dy int
}

func main() {
 	imageSize := imageSize{
		dx: 5,
		dy: 6,
	}

	imageData := Pic(imageSize.dx, imageSize.dy)
	
	pic.Show(Pic)
	rect := image.Rect(0, 0, imageSize.dx, imageSize.dy)
	img := image.NewGray(rect)
	for y := 0; y < imageSize.dy; y++ {
		for x := 0; x < imageSize.dx; x++ {
		img.Set(x, y, color.Gray{Y: imageData[y][x]})
		}
	}
	
	pic.ShowImage(img)
}
