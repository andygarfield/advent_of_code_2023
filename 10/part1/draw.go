package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// screen represents every pixel on the screen in x, y, rgba format
type screen struct {
	pixels     [980][980]color.RGBA
	area       area
	traversers []*traverser
	raster     *canvas.Raster
	window     fyne.Window
}

func (screen) ColorModel() color.Model {
	return color.RGBAModel
}

func (screen) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{980, 980}}
}

func (s *screen) At(x, y int) color.Color {
	return s.pixels[y][x]
}

func (s *screen) refresh() {
	inputArea := s.area
	for x := 0; x < len(s.pixels[0]); x++ {
		for y := 0; y < len(s.pixels); y++ {
			s.pixels[x][y] = color.RGBA{255, 255, 255, 255}
		}
	}
	for y := range inputArea {
		for x := range inputArea[y] {
			if inputArea[y][x] == '|' {
				s.drawPipe(x, y, north, south)
			} else if inputArea[y][x] == '-' {
				s.drawPipe(x, y, east, west)
			} else if inputArea[y][x] == '7' {
				s.drawPipe(x, y, west, south)
			} else if inputArea[y][x] == 'F' {
				s.drawPipe(x, y, east, south)
			} else if inputArea[y][x] == 'L' {
				s.drawPipe(x, y, east, north)
			} else if inputArea[y][x] == 'J' {
				s.drawPipe(x, y, west, north)
			}
		}
	}

	for _, t := range s.traversers {
		for x := t.current.x*7 + 1; x < t.current.x*7+6; x++ {
			for y := t.current.y*7 + 1; y < t.current.y*7+6; y++ {
				s.pixels[y][x] = color.RGBA{255, 0, 0, 127}
			}
		}
	}

	s.window.SetContent(s.raster)
}

// drawLine draws a 4 pixel line starting from x,y and going down if vertical,
// and to the right if horizontal
func (s *screen) drawPipe(x, y int, direction1, direction2 direction) {
	s.drawPipeSection(x, y, direction1)
	s.drawPipeSection(x, y, direction2)
}

func (s *screen) drawPipeSection(x, y int, direction direction) {
	if direction == north {
		for thisY := y * 7; thisY < y*7+4; thisY++ {
			s.pixels[thisY][x*7+3] = color.RGBA{0, 0, 0, 255}
		}
	} else if direction == south {
		for thisY := y*7 + 3; thisY < y*7+7; thisY++ {
			s.pixels[thisY][x*7+3] = color.RGBA{0, 0, 0, 255}
		}
	} else if direction == west {
		for thisX := x * 7; thisX < x*7+4; thisX++ {
			s.pixels[y*7+3][thisX] = color.RGBA{0, 0, 0, 255}
		}
	} else {
		for thisX := x*7 + 3; thisX < x*7+7; thisX++ {
			s.pixels[y*7+3][thisX] = color.RGBA{0, 0, 0, 255}
		}
	}
}
